package main

import (
	"log"

	"github.com/itsDavidev/TwitterGo/db"
	"github.com/itsDavidev/TwitterGo/handlers"
)

func main() {
	if db.IsConnection() == false {
		log.Fatal("Connection to mongoDB failed")
		return
	}
	handlers.HandlersRoutes()

}
