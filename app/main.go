package main

import (
	routes "github.com/ronnachate/inventory-api-go/api/route"

	initializers "github.com/ronnachate/inventory-api-go/initializer"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	initializers.SetupDatabase(&config)
}

func main() {
	r := routes.SetupRouter()

	r.Run()
}
