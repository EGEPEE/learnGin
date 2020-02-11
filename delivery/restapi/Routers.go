package restapi

import (
	"log"

	"github.com/EGEPEE/learnGin/controllers"
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
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

	// User
	e.GET("/get_all_account", controllers.Mobile(auth), controllers.GetAllAcount)
	e.POST("/check_phonenumber", controllers.Mobile(auth), controllers.CheckPhone)
	e.POST("/delete_account", controllers.Mobile(auth), controllers.DeleteAccount)
	e.POST("/check_pin", controllers.Mobile(auth), controllers.CheckPin)
	e.POST("/set_pin", controllers.Mobile(auth), controllers.SetPin)
	e.POST("/register", controllers.Mobile(auth), controllers.Register)
	e.POST("/forgot_pin", controllers.Mobile(auth), controllers.ForgotPin)

	return e
}
