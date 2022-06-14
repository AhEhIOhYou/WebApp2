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

	r.LoadHTMLGlob("app/view/templates/*.tmpl")

	r.GET("/", controller.Home)
	r.GET("/staff", controller.GetStaff)
}
