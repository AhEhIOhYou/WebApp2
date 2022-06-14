package controller

import (
	"awesomeProject1/app/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetStaff(c *gin.Context) {

	staff, err := model.GetAllStaff()

	if err != nil {
		c.AbortWithError(400, err)
		log.Fatal(err)
		return
	}

	c.HTML(http.StatusOK, "staff.tmpl", gin.H{
		"title": "Staff",
		"data":  staff,
	})
}
