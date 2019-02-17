package member

import (
	"go-member/internal/app"

	"github.com/globalsign/mgo"
)

func insertMemberDB(member Member) error {
	session, err := app.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	db := session.DB(databaseMember)

	if err := db.C("member").Insert(member); err != nil {
		return err
	}
	return nil
}

func applyMongoIncrement(change mgo.Change, doc *IncrementIndex) error {
	session, err := app.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	db := session.DB(databaseMember)

	_, err = db.C("increment_index").Find(nil).Apply(change, &doc)
	if err != nil {
		return err
	}
	return nil
}

func createMongoIncrementCollection(doc IncrementIndex) error {
	session, err := app.GetMongoSession()
	if err != nil {
		return err
	}
	defer session.Close()
	db := session.DB(databaseMember)

	if err := db.C("increment_index").Insert(doc); err != nil {
		return err
	}
	return nil
}
