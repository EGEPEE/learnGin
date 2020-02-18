package controllers

import (
	"fmt"
	"os"
	"time"

	"github.com/EGEPEE/learnGin/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type Role int

const (
	ADMIN        Role = 0x1
	FINANCE      Role = 0x1 << 1
	MOBILE       Role = 0x1 << 2
	SYSTEM_ADMIN Role = 0x1 << 2
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func HelloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"Username": user.(*models.User).Username,
		"role":     user.(*models.User).Role,
		"text":     "Hello World.",
	})
}

func NewAuth() (authMiddleware *jwt.GinJWTMiddleware, err error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(os.Getenv("SCRT_KEY")),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			u := models.NaiveDatastore[claims[identityKey].(string)]
			return &models.User{
				Username: claims[identityKey].(string),
				Role:     u.Role,
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			request := login{Username: c.PostForm("username"), Password: c.PostForm("password")}
			if err := c.ShouldBind(&request); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			u := models.NaiveDatastore[request.Username]
			if u.Password != request.Password {
				return nil, jwt.ErrFailedAuthentication
			}

			return &u, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.Username == "mobile" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
}

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, _ := c.Get(identityKey)
		fmt.Println(user.(*models.User).Role)
	}
}
