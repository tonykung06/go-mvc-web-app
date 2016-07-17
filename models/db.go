package models

import (
	"gopkg.in/mgo.v2"
)

func GetDbConnection() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}

func getCollection(session *mgo.Session, collection string) *mgo.Collection {
	return session.DB("go-mvc-web-app").C(collection)
}
