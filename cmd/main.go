package main

import (
	routes "github.com/ronnachate/inventory-api-go/api/route"
)

func main() {
	r := routes.SetupRouter()

	r.Run()
}
