package main

import (
	"awesomeProject1/app/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()

	router.Static("assets/js/", "assets/js/")
	router.Static("assets/css/", "assets/css/")

	routes(router)
	err := router.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *gin.Engine) {

	r.LoadHTMLGlob("app/view/*.html")

	r.GET("/", controller.Home)

	//r.GET("/staff", controller.GetStaff)
	//r.GET("/staff:id", controller.GetStaff)
	//
	//r.GET("/API-test", controller.Test)

	r.GET("/staff", controller.GetStaff)
	r.Any("/staff/create", controller.CreateStaff)

}
