package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and returns it.
// We can now use this function to instantiate and test the router outside the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	return r
}

func main() {
	// The router is now formed by calling the 'newRouter'constructor function defined above.
	r := newRouter()
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
