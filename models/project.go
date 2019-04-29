package models

import (
	"github.com/globalsign/mgo/bson"
)

const Collection = "project"

type Project struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name"`
}

type ProjectModel struct{}

func (m ProjectModel) Create(project Project) (err error) {
	collection := DB.C(Collection)
	err = collection.Insert(project)
	return err
}
