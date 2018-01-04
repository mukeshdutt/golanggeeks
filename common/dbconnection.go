package common

import (
	"gopkg.in/mgo.v2"
)

// Session is
var Session *mgo.Session

// DB is
const DB = "golanggeeks"

// Initialize is
func Initialize() *mgo.Session {
	Session, err := mgo.Dial("localhost")
	if err != nil {
		panic("error occurred")
	}
	return Session
}
