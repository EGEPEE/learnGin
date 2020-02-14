package restapi

import (
	"log"

	"github.com/EGEPEE/learnGin/controllers"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	auth, err := controllers.NewAuth()
	if err != nil {
		log.Fatal(err)
	}

	e := gin.New()

	e.Use(jwt.ErrorHandler)
	e.POST("/login", auth.Authenticate)
	e.POST("/auth/refresh_token", auth.RefreshToken)

	// USER
	user := e.Group("/api/usr_userapi")
	{
		user.GET("/get_all_account", controllers.Mobile(auth), controllers.GetAllAcount)
		user.POST("/check_phonenumber", controllers.Mobile(auth), controllers.CheckPhone)
		user.POST("/delete_account", controllers.Mobile(auth), controllers.DeleteAccount)
		user.POST("/check_pin", controllers.Mobile(auth), controllers.CheckPin)
		user.POST("/set_pin", controllers.Mobile(auth), controllers.SetPin)
		user.POST("/register", controllers.Mobile(auth), controllers.Register)
		user.POST("/forgot_pin", controllers.Mobile(auth), controllers.ForgotPin)
	}

	// PICKER

	// BANK SAMPAH
	return e
}
