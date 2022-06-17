package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Home",
	})
}

func Modal(c *gin.Context) {
	c.HTML(http.StatusOK, "testmodal.html", gin.H{
		"title": "AaaLlllol",
	})
}

func Modal2(c *gin.Context) {
	c.HTML(http.StatusOK, "testmodal.html", gin.H{
		"title":   "AaaLlllol",
		"success": "true",
	})
}
