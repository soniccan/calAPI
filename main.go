package main

import (
	"calAPI/routes"
)

func main() {

	router := routes.NewRoutes()
	router.Run()
}
