package config

import (
	"gopkg.in/mgo.v2"
)

func GetConnect() (*mgo.Database, error) {
	host := "mongodb://localhost:27017"
	dbName := "demo_session88"
	session, err := mgo.Dial(host)
	if err != nil {
		return nil, err
	}
	db := session.DB(dbName)
	return db, nil
}
