package member

import (
	"strconv"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// AddErrorDetail is a function for appending error to a response
func (responseError *Error) AddErrorDetail(errorDetail ErrorDetail) []ErrorDetail {
	responseError.Details = append(responseError.Details, errorDetail)
	return responseError.Details
}

// genCustomerID is a function for generating customer ID using mongo increment feature
func genCustomerID() (string, error) {
	doc := IncrementIndex{}
	change := mgo.Change{
		Update: bson.M{"$inc": bson.M{
			"customer_id": 1,
		}},
		ReturnNew: true,
	}

	err := applyMongoIncrement(change, &doc)
	if err != nil {
		if err.Error() == "not found" {
			doc.CustomerID = 1
			if err := createMongoIncrementCollection(doc); err != nil {
				return "", err
			}
		} else {
			return "", err
		}
	}
	return strconv.Itoa(doc.CustomerID), nil
}
