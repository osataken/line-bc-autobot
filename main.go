package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/message/relay", MessageRelayHandler)

	// Bind to a port and pass our router in
	http.ListenAndServe(":80", r)
}

func MessageRelayHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}