package restapi

import (
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	auth, err := NewAuth()

	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	return r
}
