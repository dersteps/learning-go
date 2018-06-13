package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	// Install with go get...
	"github.com/gorilla/mux"
)

func addRoute(router *mux.Router, path string, f http.HandlerFunc, methods ...string) {
	fmt.Printf("Adding route '%s' [%s] -> %s\n", path, strings.Join(methods, ", "), reflect.TypeOf(f).Name())
	router.HandleFunc(path, f).Methods(methods...)
}

func handleRootRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "List of endpoints and stuff")
}

func main() {
	const addr string = ":9991"
	fmt.Println("Static REST API is starting")

	// Create new router
	router := mux.NewRouter()

	fmt.Println("Setting up routes...")

	// Set up routes
	addRoute(router, "/", handleRootRequest, "GET")
	fmt.Printf("Serving static API at %s\n", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		fmt.Println("Whoops, that failed horribly!")
		panic(err)
	}
}
