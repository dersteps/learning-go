package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Credentials struct {
	name     string
	password string
}

var credentials = []Credentials{}

func handleRequests(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Alright, welcome\n"))
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

// Middleware for basic auth
func basicAuth(wrapped http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("WWW-Authenticate", `Basic realm="Restricted area"`)

		s := strings.SplitN(req.Header.Get("Authorization"), " ", 2)
		if len(s) != 2 {
			http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		b, err := base64.StdEncoding.DecodeString(s[1])
		if err != nil {
			http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		pair := strings.SplitN(string(b), ":", 2)
		if len(pair) != 2 {
			http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		for _, cred := range credentials {
			if cred.name == pair[0] && cred.password == pair[1] {
				fmt.Printf("User accepted: %s\n", pair[0])
				wrapped.ServeHTTP(res, req)
				return
			}
		}

		http.Error(res, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)

	}
}

func main() {

	credentials = append(credentials, Credentials{"steps", "steps"})
	credentials = append(credentials, Credentials{"user", "user"})
	credentials = append(credentials, Credentials{"root", "toor"})

	router := mux.NewRouter()
	fmt.Printf("Setting up route /basic...\n")
	router.HandleFunc("/basic", use(handleRequests, basicAuth))
	addr := ":10101"
	fmt.Printf("Starting basic authentication server on %s\n", addr)

	http.ListenAndServe(addr, router)

}
