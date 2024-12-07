package main

import (
	"github.com/DenysNykoriak/go-api/initialization"
	"github.com/gin-gonic/gin"
)

func init() {
	initialization.LoadEnv()
	initialization.ConnectPostgres()

	initialization.SyncPostgres()
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run()

}
