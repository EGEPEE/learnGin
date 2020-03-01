package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/EGEPEE/learnGin/delivery/restapi"
	"github.com/EGEPEE/learnGin/models"
)

func main() {
	// err := inject.LoadCasbinPolicyData()
	// if err != nil {
	// 	panic("Kesalahan memuat data kebijakan casbin: " + err.Error())
	// }

	routersInit := restapi.InitRouter()
	readTimeout := models.ServerSetting.ReadTimeout
	writeTimeout := models.ServerSetting.WriteTimeout

	endPoint := fmt.Sprintf(":%d", models.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	_ = server.ListenAndServe()

}
