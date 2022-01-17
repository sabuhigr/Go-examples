package main

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/brandenc40/romannumeral"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlPathElements := strings.Split(r.URL.Path, "/")
		if urlPathElements[1] == "roman_number" {
			number, err := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "%s\n", "Request variable is not integer")
				w.Write([]byte("Usage : http://127.0.0.1:8000/roman_number/{integer}"))
				return
			}
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("Requested number is 0 or bigger than 10"))
			} else {
				val, ok := romannumeral.IntToString(number)
				if ok != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				fmt.Fprintf(w, "%q", html.EscapeString(val))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Status Code : [400], Bad request"))
		}
	})

	s := &http.Server{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
