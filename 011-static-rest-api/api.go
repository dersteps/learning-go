package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	// Install with go get...
	"github.com/gorilla/mux"
)

/**
 *  Struct that represents a book.
 */
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

/*
	The (b Book) part tells Go that this func is related to the Book type,
	i.e. is a method of a book.
*/
func (b Book) toJSON() string {
	return toJSON(b)
}

func toJSON(x interface{}) string {
	bytes, err := json.Marshal(x)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	return string(bytes)
}

func readBooks(file string) []Book {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Unable to read books from %s: %s", file, err.Error())
		panic(err)
	}

	fmt.Println(string(data))

	// Convert to slice of Book
	var books []Book
	json.Unmarshal(data, &books)

	return books
}

func addRoute(router *mux.Router, path string, f http.HandlerFunc, methods ...string) {
	fmt.Printf("Adding route '%s' [%s] -> %s\n", path, strings.Join(methods, ", "), reflect.TypeOf(f).Name())
	router.HandleFunc(path, f).Methods(methods...)
}

func handleRootRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "List of endpoints and stuff")
}

func handleBooksRequest(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(toJSON(books)))
	res.Write([]byte("\n"))
}

func handleBookRequest(res http.ResponseWriter, req *http.Request) {
	// Attempt to find the book
	vars := mux.Vars(req)
	reqID := vars["id"]

	for _, book := range books {
		fmt.Printf("Testing book id '%d' against requested ID '%s'\n", book.ID, reqID)
		d, err := strconv.Atoi(reqID)
		if err != nil {
			http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			//res.WriteHeader(http.StatusBadRequest)
			//res.Write([]byte("Invalid book id"))
			return
		}

		if book.ID == d {
			res.Write([]byte(toJSON(book)))
			res.Write([]byte("\n"))
			return
		}
	}
	//res.WriteHeader(http.StatusNotFound)
	//res.Write([]byte("[]\n"))
	http.Error(res, http.StatusText(http.StatusNotFound), http.StatusNotFound)

}

func main() {
	const addr string = ":9991"
	fmt.Println("Static REST API is starting")

	// Create new router
	router := mux.NewRouter()

	fmt.Println("Setting up routes...")

	// Set up routes
	addRoute(router, "/", handleRootRequest, "GET")

	addRoute(router, "/books", handleBooksRequest, "GET")

	addRoute(router, "/books/{id}", handleBookRequest, "GET")

	// Get content
	books = readBooks("books.json")

	fmt.Println(toJSON(books))

	fmt.Printf("Serving static API at %s\n", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		fmt.Println("Whoops, that failed horribly!")
		panic(err)
	}
}
