package main

import (
	"fmt"
	"net/http"
)

func echo(res http.ResponseWriter, req *http.Request) {
	written, err := fmt.Fprintf(res, "You have requested %s", req.URL.Path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Written % 4d bytes to response stream\n", written)

}

func main() {
	fmt.Println("Setting up routes")
	http.HandleFunc("/", echo)
	fmt.Println("Serving...")
	err := http.ListenAndServe(":9900", nil)
	if err != nil {
		panic(err)
	}

}
