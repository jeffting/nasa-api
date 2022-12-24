package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeffting/nasa-api/internal/clients"
)

// Set up our server
func Serve() {
	router := mux.NewRouter()

	clients := clients.InitializeClients()
	routes(router, clients)
	fmt.Println("listening on port :8081")
	http.ListenAndServe(":8081", router)
}
