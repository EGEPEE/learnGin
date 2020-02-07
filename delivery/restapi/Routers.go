package restapi

import (
	"github.com/EGEPEE/learnGin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("/api/usr_userapi")
	{
		user.GET("/get_user", controllers.GetUser)
		user.POST("/check_phonenumber", controllers.CheckPhone)
		user.POST("/register", controllers.Register)
	}

	return r
}
