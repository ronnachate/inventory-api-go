package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/api/middleware"
	"github.com/ronnachate/inventory-api-go/api/route"
	infrastructure "github.com/ronnachate/inventory-api-go/infrastructure"
)

func main() {
	config, err := infrastructure.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	infrastructure.SetupDatabase(&config)
	defer infrastructure.CloseDBConnection()

	contextTimeout := time.Duration(config.ContextTimeout) * time.Second

	gin := gin.Default()
	infrastructure.Logger = infrastructure.SetupLogger(config)

	gin.Use(middleware.ApplicationKeyMiddleware(config))

	route.SetupRouter(infrastructure.DB, contextTimeout, gin)

	gin.Run()
}
