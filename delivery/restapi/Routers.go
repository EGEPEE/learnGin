package restapi

import (
	"github.com/EGEPEE/learnGin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	user := r.Group("/api/usr_userapi")
	{
		user.GET("/get_all_account", controllers.GetAllAcount)
		user.POST("/check_phonenumber", controllers.CheckPhone)
		user.POST("/delete_account", controllers.DeleteAccount)
		user.POST("/check_pin", controllers.CheckPin)
		user.POST("/set_pin", controllers.SetPin)
		user.POST("/register", controllers.Register)
	}

	return r
}
