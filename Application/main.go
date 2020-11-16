package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

// Declaring the router
var router = mux.NewRouter()

// Main function
func main() {
	Routes()
	http.ListenAndServe(":8000", router)
}
