package db

import (
	"log"
	"github.com/osataken/line-bc-autobot/message"
)

func SaveReceivedMessage(receivedMessage message.Result) {
	session := GetSession()
	defer session.Close()

	c := session.DB("heroku_7x3w65x3").C("received_messages")
	err := c.Insert(&receivedMessage)
	if err != nil {
		log.Fatal(err)
	}
}