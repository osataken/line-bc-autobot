package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"os"
)

func main() {
	println("Starting up line bc autobot!!!")

	r := mux.NewRouter()
	r.HandleFunc("/", DefaultHandler)
	r.HandleFunc("/message/relay", MessageRelayHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		println("Error occurred: " + err.Error())
		panic(err)
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func MessageRelayHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}