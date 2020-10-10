package controllers

import (
	"fmt"
	"net/http"
)

// ContextRootController handles the root endpoint of the API.
func ContextRootController(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}
