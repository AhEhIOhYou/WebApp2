package controller

import (
	"awesomeProject1/app/model"
	"awesomeProject1/app/server"
	"fmt"
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

		fmt.Println(staff)

		c.HTML(200, "staff.html", gin.H{
			"title":    "Home",
			"dataRows": staff,
		})
		return
	}

	staff, err := model.GetAllStaff(db)

	if err != nil {
		c.AbortWithError(400, err)
		log.Fatal(err)
		return
	}

	c.HTML(200, "staff.html", gin.H{
		"title":    "Home",
		"dataRows": staff,
	})
}
