package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to this life-changing API. It should hot reload!")
	})

	fmt.Println("Server listening!")
	http.ListenAndServe(":8080", r)
}
