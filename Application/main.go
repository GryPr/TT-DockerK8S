package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Home endpoint
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// Main function
func main() {
	// Defining the router
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)
	http.ListenAndServe(":8000", router)
}
