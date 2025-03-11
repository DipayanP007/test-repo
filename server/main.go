package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	log.Println("Golang API")
	log.Println("Listening on localhost:4000")
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to apis in golang</h1>"))
}
