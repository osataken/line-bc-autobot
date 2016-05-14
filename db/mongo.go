package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)

var mongoSession *mgo.Session

func InitDB() {
	var err error

	url := os.Getenv("MONGODB_URI")
	session, err := mgo.Dial(url)
	//session, err := mgo.Dial("mongodb://192.168.99.100:27017/heroku_7x3w65x3")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	mongoSession = session

	fmt.Println("Init MongoDB Connection Completed.")
}

func GetSession() *mgo.Session {
	return mongoSession.Copy()
}