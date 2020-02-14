package controllers

import (
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type login struct {
	NoTelepon string `form:"no_telepon json:"no_telepon" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

var identityKey = os.Getenv("IDNTY")

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	user, _ := c.Get(identityKey)
	c.JSON(200, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func NewAuth() {
	authMiddleware, err := jwt.New(
		&jwt.GinJWTMiddleware{
			Realm:      "test zone",
			Key:        []byte(os.Getenv("SCRT_KEY")),
			Timeout:    time.Hour,
			MaxRefresh: time.Hour,
			IdentityKey: os.Getenv("IDNTY"),
			PayloadFunc: func(data interface{}) jwt.MapClaims {
				if v, ok := data.(*User); ok {
					return jwt.MapClaims{
						identityKey: v.UserName,
					}
				}
			},
			IdentityHandler: func(c *gin.Context) interface{} {
				claims: jwt.ExtractClaims
			}
		})
}
