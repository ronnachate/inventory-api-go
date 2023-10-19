package main

import (
	"github.com/ronnachate/inventory-api-go/routes"
)

func main() {
	r := routes.SetupRouter()

	r.Run()
}
