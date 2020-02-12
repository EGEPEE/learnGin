package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/EGEPEE/learnGin/models"

	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
)

func NewAuth() (jwt.Auth, error) {
	return jwt.New(jwt.Auth{
		SecretKey: []byte(os.Getenv("TKN_JWT")),
		Authenticator: func(c *gin.Context) (jwt.MapClaims, error) {
			type req struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			request := req{Username: c.PostForm("username"), Password: c.PostForm("password")}
			if err := c.ShouldBind(&request); err != nil {
				fmt.Println("error cek uname & pass")
				return nil, jwt.ErrorAuthenticationFailed
			}

			u := models.NaiveDatastore[request.Username] // change here fetching from read datastore
			if u.Password != request.Password {
				c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Username atau password salah."})

				return nil, jwt.ErrorAuthenticationFailed
			}

			return jwt.MapClaims{
				"username": u.Username,
				"role":     u.Role,
			}, nil
		},
		UserFetcher: func(c *gin.Context, claims jwt.MapClaims) (interface{}, error) {
			username, ok := claims["username"].(string)
			if !ok {
				return nil, nil
			}
			u, ok := models.NaiveDatastore[username]
			if !ok {
				return nil, nil
			}
			return u, nil
		},
	})
}

func Admin(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsAdmin()
	})
}

func Finance(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsFinance()
	})
}

func Mobile(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsMobile()
	})
}

func SystemAdmin(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsSystemAdmin()
	})
}

func role(claims jwt.MapClaims) models.Role {
	return models.Role(claims["role"].(float64))
}
