package main

import (
	"awesomeProject1/app/controller"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	routes(router)
	err := router.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *gin.Engine) {
	//что следует выполнять при входящих запросах указанного типа и по указанному адресу
	r.LoadHTMLGlob("app/view/templates/*.tmpl")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	r.GET("/", controller.IndexPage)
	r.GET("/staff", controller.GetStaff)
}
