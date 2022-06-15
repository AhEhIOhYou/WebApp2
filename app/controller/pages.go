package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Home",
		"test":  "aa - home",
		"text":  "some text- home",
	})
}
