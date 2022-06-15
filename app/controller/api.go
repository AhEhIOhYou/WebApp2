package controller

import (
	"awesomeProject1/app/model"
	"awesomeProject1/app/server"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	db := server.Connect()
	staff, _ := model.GetStaffById(db, 1)
	c.JSON(200, staff)
}
