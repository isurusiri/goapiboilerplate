package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/isurusiri/goapiboilerplate/controllers"
)

// App holds the route instance of the API and
// exposes a set of function to initialize and
// interact with it.
type App struct {
	router *mux.Router
}

// Initialize , initializes the API routes.
func (a *App) Initialize() {
	a.router = mux.NewRouter().StrictSlash(true)

	a.router.HandleFunc("/", controllers.ContextRootController)
}

// Run the API.
func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(port, a.router))
}
