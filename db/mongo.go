package db

import "gopkg.in/mgo.v2"

var session mgo.Session

func InitDB() {
	var err error
	session, err := mgo.Dial("mongodb://line-bc-autobot:ilovedata@ds023042.mlab.com:23042/heroku_7x3w65x3")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
}

func GetSession() *mgo.Session {
	return session.Copy()
}