package main

import (
	"calAPI/routes"
	"log"
)

func main() {

	router := routes.NewRoutes()
	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}
