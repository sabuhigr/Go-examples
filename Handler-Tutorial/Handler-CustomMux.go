package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

type CustomServeMux struct{} //Our custom ServeMux


//If any type implement ServeHTTP, it can be Mux
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		fmt.Println(r.URL.User)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main() {

	mux := &CustomServeMux{}
	http.ListenAndServe(":8000", mux)
}
