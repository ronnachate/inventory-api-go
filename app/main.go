package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ronnachate/inventory-api-go/api/route"

	initializer "github.com/ronnachate/inventory-api-go/initializer"
)

func main() {
	config, err := initializer.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	initializer.SetupDatabase(&config)
	defer initializer.CloseDBConnection()

	contextTimeout := time.Duration(config.ContextTimeout) * time.Second

	gin := gin.Default()
	route.SetupRouter(initializer.DB, contextTimeout, gin)

	gin.Run()
}
