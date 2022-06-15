package controller

import (
	"awesomeProject1/app/model"
	"awesomeProject1/app/server"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetStaff(c *gin.Context) {

	id := c.Query("id")
	db := server.Connect()

	if id != "" {
		id, _ := strconv.Atoi(id)
		staff, err := model.GetStaffById(db, id)

		if err != nil || staff.Id == 0 {
			c.JSON(200, model.AppError{
				Code:    "NotFound",
				Details: "Data not found in database",
			})
			log.Println(err)
			return
		}

		c.JSON(200, staff)
		return
	}

	staff, err := model.GetAllStaff(db)

	if err != nil {
		c.AbortWithError(400, err)
		log.Fatal(err)
		return
	}

	c.JSON(200, staff)
}

func Create(c *gin.Context) {
	c.HTML(200, "create.tmpl", gin.H{
		"title": "Create",
	})
}
