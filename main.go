package main

import (
	"fiber-go-crud/api"
	"fiber-go-crud/database"
	"log"
)

func main() {
	if err := database.ConnectionDB(); err != nil {
		log.Fatal(err)
	}

	api.Init()
}
