package main

import (
	"net/http"
	"os"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/osataken/line-bc-autobot/message"
	"github.com/osataken/line-bc-autobot/line"
	"github.com/osataken/line-bc-autobot/db"
	"github.com/osataken/line-bc-autobot/template"
)

func main() {
	fmt.Println("Starting up line bc autobot!!!")
	http.HandleFunc("/", DefaultHandler)
	http.HandleFunc("/message/relay", MessageRelayHandler)
	http.HandleFunc("/message/send", MessageSendHandler)
	http.HandleFunc("/registration", RegistrationFormHandler)
	http.HandleFunc("/registration/save", RegistrationFormSaveHandler)

	db.InitDB()

	err := http.ListenAndServe(":" + os.Getenv("PORT"), nil)
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
		fmt.Println(string(body))
		err := json.Unmarshal(body, &receivedMessage)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(400)
			w.Write([]byte("bad request!"))
			return
		}

		if receivedMessage.Result != nil && len(receivedMessage.Result) > 0 {
			for _, result := range receivedMessage.Result {
				handleRelayMessage(w, result)
			}
		}
	}
}

func writeResponse(w http.ResponseWriter, text string) {
	w.Write([]byte(text))
}

func RegistrationFormHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, template.GetRegistrationForm())
}

func RegistrationFormSaveHandler(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, "Thank you for registration!!")

	r.ParseForm()

	for key, value := range r.Form {
		fmt.Println("Key:", key, " Value:", value)
	}
}

func handleRelayMessage(w http.ResponseWriter, result message.Result) {

	if result.From != "" {
		db.SaveReceivedMessage(result)

		text := fmt.Sprintf("Received message: %v", result.Content.Text)
		w.Write([]byte(text + "\n"))

		mid := result.Content.From

		_, err := sendOAText(mid, text);
		if err != nil {
			text := fmt.Sprintf("SendErr: %v", err.Error())
			w.Write([]byte(text))
			return
		}
	}
}


func MessageSendHandler(w http.ResponseWriter, r *http.Request) {
	mid := r.URL.Query().Get("mid")
	content := r.URL.Query().Get("content")

	messageId, err := sendOAText(mid, content);
	if err != nil {
		text := fmt.Sprintf("Err: %v", err.Error())
		w.Write([]byte(text))
		return
	}

	text := fmt.Sprintf("messageId: %v", messageId)
	w.Write([]byte(text))
}

func sendOAText(mid, content string) (string, error) {
	fmt.Println("SendOATo:", mid, " content:", content)

	req := &message.EventsRequest{
		To:        []string{ mid },
		ToChannel: 1383378250,
		EventType: "138311608800106203",
		Content: &message.Content{
			ContentType: 1,
			ToType:      1,
			Text:        content,
		},
	}

	rsp := &message.EventsResponse{}

	accessToken := os.Getenv("OA_CHANNEL_TOKEN")

	err := line.CallLineApi("https://api.line.me/v1/", "events", accessToken, req, rsp)
	if err != nil {
		return "", err
	}

	fmt.Println("Success with messageId:", rsp.MessageId)
	return rsp.MessageId, nil

}