package results

import "github.com/globalsign/mgo/bson"

type ProjectResult struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
}
