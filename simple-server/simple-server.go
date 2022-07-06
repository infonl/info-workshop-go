package simpleServer

import (
	"github.com/gin-gonic/gin"
	helloWorld "info-workshop-go/hello-world"
	"net/http"
)

func Server() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": helloWorld.GetHelloWorld(),
		})
	})
	r.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{
			"message": helloWorld.GetHelloName(name),
		})
	})
	r.Run()
}
