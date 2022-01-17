package main

import (
	"io"
	"log"
	"net/http"
)
//----Alternative to h1 := func(w http.ResponseWriter, _ *http.Request)
//--- in Golang, |h1 := func(a,b string)| equal to |func h1(a,b string)|

// func h1(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello, World"))
// }

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	//------Alternative to above function and below HandleFunc---------

	// http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
	// 	io.WriteString(w, "Hello from a HandleFunc #1!\n")
	// })
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8111", nil))
}
