package main

import (
	"net/http"
	"os"
	"fmt"
)

func main() {
	fmt.Println("Starting up line bc autobot!!!")
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/message/relay", MessageRelayHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		fmt.Println("Error occurred: " + err.Error())
		panic(err)
	}
}

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}

func MessageRelayHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}