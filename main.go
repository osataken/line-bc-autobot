package main

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/osataken/line-bc-autobot/message"
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
	receivedMessage := message.Receive{}


	if body, _ := ioutil.ReadAll(r.Body); len(body) > 0 {
		err := json.Unmarshal(body, &receivedMessage)
		if err != nil {
			w.Write([]byte("bad request!"))
		}

		text := fmt.Sprintf("Received message: %v", receivedMessage.Result[0].Content.Text)
		w.Write([]byte(text))
	}
}