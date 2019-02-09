package app

import (
	"time"

	"github.com/globalsign/mgo"
)

func GetMongoSession() (*mgo.Database, error) {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{CFG.MongoDB.MongoDBHost},
		Timeout:  CFG.MongoDB.Timeout * time.Second,
		Database: CFG.MongoDB.AuthDatabase,
		Username: CFG.MongoDB.Username,
		Password: CFG.MongoDB.Password,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	database := session.DB("go-member")
	return database, err
}
