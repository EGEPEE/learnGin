package restapi

import (
	"log"

	"github.com/EGEPEE/learnGin/controllers"
	jwt "github.com/appleboy/gin-jwt"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	authMiddleware, err := controllers.NewAuth()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/login", authMiddleware.LoginHandler)
	// Refresh time can be longer than token timeout
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	user := r.Group("/user")
	{
		user.GET("/hello", controllers.Admin(), controllers.HelloHandler)
	}

	return r
}
