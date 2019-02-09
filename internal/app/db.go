package app

import (
	"time"

	"github.com/globalsign/mgo"
)

func GetMongoSession() (*mgo.Database, error) {
	MongoDBHost := "localhost:27017"
	AuthDatabase := "admin"
	AuthUserName := "admin"
	AuthPassword := "passw0rd"

	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHost},
		Timeout:  10 * time.Second,
		Database: AuthDatabase,
		Username: AuthUserName,
		Password: AuthPassword,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	database := session.DB("go-member")
	return database, err
}
