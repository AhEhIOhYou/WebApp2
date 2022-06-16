package controller

import (
	"awesomeProject1/app/model"
	"awesomeProject1/app/server"
	"github.com/gin-gonic/gin"
)

func CreateStaff(c *gin.Context) {

	if c.Request.Method == "POST" {
		c.Request.ParseForm()
		params := c.Request.Form
		name := params.Get("name")
		if model.CreateStaff(server.Connect(), name) != nil {
			c.HTML(200, "create.tmpl", gin.H{
				"title": "CreateStaff",
			})
		}
	}

	fields, _ := model.GetFields(server.Connect(), "aaa")

	c.HTML(200, "create.html", gin.H{
		"title":  "CreateStaff",
		"fields": fields,
	})
}
