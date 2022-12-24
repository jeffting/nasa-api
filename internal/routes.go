package internal

import (
	"github.com/jeffting/nasa-api/internal/clients"
	"github.com/jeffting/nasa-api/internal/handlers"

	"github.com/gorilla/mux"
)

func routes(router *mux.Router, clients clients.Clients) {
	// insert routes here
	// gets images for the last 10 days
	router.Handle("/v1/rover/images", handlers.GetImagesHandler(clients))
	// gets images for a specified day
	router.Handle("/v1/rover/images/day", handlers.GetImagesDayHandler(clients))
}
