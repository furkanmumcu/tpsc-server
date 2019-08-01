package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		fmt.Println("hello called")
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		fmt.Println("called")
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"user": name})
	})

	r.Run(":9191")
}
