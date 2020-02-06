package restapi

import (
	"github.com/EGEPEE/learnGin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/usr_userapi")
	{
		v1.GET("/get_user", controllers.GetUser)
	}

	return r
}
