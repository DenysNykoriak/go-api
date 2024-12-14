package main

import (
	app_auth "github.com/DenysNykoriak/go-api/app/auth"
	"github.com/DenysNykoriak/go-api/core"
	"github.com/gin-gonic/gin"
)

func init() {
	core.LoadEnv()
	core.ConnectPostgres()

	core.SyncPostgres()
}

func main() {
	r := gin.Default()

	app_auth.InitializeRoutes(r)

	r.Run()

}
