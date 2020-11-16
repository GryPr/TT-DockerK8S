package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var integerArray [50]int
var intArrayLocation int = 0

type intRequest struct {
	Integer int `json:"integer"`
}

// Routes lists the endpoints of the API
func Routes() {
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/intstorage", intStore)
}

// Home endpoint
func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func intStore(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var number intRequest
	err := decoder.Decode(&number)
	if err != nil {
		panic(err)
	}
	integerArray[intArrayLocation] = number.Integer
	intArrayLocation++
	json.NewEncoder(w).Encode(integerArray)
}
