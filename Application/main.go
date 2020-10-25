package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Creating an event struct
type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

// Home endpoint
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// Main function
func main() {
	// Defining the router
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
}
