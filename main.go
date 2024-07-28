package main

import (
	"example.com/m/v2/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEndVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
