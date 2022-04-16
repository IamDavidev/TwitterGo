package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandlersRoutes() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handlerUserAccess := cors.AllowAll().Handler(router)

	addr := ":" + PORT

	log.Fatal(http.ListenAndServe(addr, handlerUserAccess))

}
