package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var count int = 0

func main() {

	HomePage := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("Hello World"))
	}

	ProductPage := func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Println(count) //for how much request you receive
		if r.Host == "127.0.0.1:8080" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("You requested from Local network -> 127.0.0.1"))
			return
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", "You requested from global network ")
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Your requested page not found"))
	}

	router := mux.NewRouter()
	router.HandleFunc("/", HomePage)
	router.HandleFunc("/products", ProductPage)
	server := &http.Server{
		Handler:      router,
		Addr:         ":8080",
		ReadTimeout:  15,
		WriteTimeout: 15,
	}

	server.ListenAndServe()
}
