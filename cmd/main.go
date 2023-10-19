package main

import (
	"github.com/ronnachate/inventory-api-go/api/routes"
)

func main() {
	r := routes.SetupRouter()

	r.Run()
}
