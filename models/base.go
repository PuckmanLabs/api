package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"reflect"
)

type Model struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

func Collection(model interface{}, DB *mgo.Database) *mgo.Collection {
	modelName := reflect.TypeOf(model).Elem().Name()
	return DB.C(modelName)
}
