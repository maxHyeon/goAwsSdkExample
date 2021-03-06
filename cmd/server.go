package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	r.HandleFunc("/sts", GetCallerIdentity)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
