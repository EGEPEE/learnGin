package controllers

import (
	"fmt"
	"time"

	"github.com/EGEPEE/learnGin/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func HelloHandler(c *gin.Context) {
	fmt.Println(c)
	// claims := jwt.ExtractClaims(c)
	// fmt.Println(claims)
	// user, _ := c.Get(identityKey)
	// fmt.Println(user)
	// c.JSON(200, gin.H{
	// 	"userID":   claims[identityKey],
	// 	"userName": user.(*models.User).Username,
	// 	// "email":    user.(*models.User).Email,
	// 	// "role":     user.(*models.User).Role,
	// 	"texting": "Hello World.",
	// })
}

func NewAuth() (authMiddleware *jwt.GinJWTMiddleware, err error) {
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
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
		// Authorizator: func(data interface{}, c *gin.Context) bool {
		// 	validatorRole := checkRole(data)

		// 	return validatorRole
		// },
		// Unauthorized: func(c *gin.Context, code int, message string) {
		// 	c.JSON(code, gin.H{
		// 		"code":    code,
		// 		"message": message,
		// 	})
		// },
		// Unauthorized: func(c *gin.Context, code int, message string) {
		// 	c.JSON(code, gin.H{
		// 		"code":    code,
		// 		"message": message,
		// 	})
		// },

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}

func Admin(claims jwt.MapClaims) bool {
	fmt.Println(claims)
	// return authMiddleware.VerifyPerm(func(claims jwt.MapClaims) bool {
	// return role(claims).IsAdmin()
	// })
	return true
}
