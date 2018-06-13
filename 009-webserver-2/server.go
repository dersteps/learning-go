package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handleBookGet(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	title := vars["title"]
	page := vars["page"]

	written, err := fmt.Fprintf(res, "Displaying book '%s' on page '%s'", title, page)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Written % 4d bytes to response stream\n", written)
}

func main() {
	r := mux.NewRouter()

	// Setup route for book example
	r.HandleFunc("/books/{title}/p/{page}", handleBookGet).Methods("GET")

	fmt.Println("Listening on port 9991...")
	err := http.ListenAndServe(":9991", r)

	if err != nil {
		panic(err)
	}
}
