package models

import (
	"fmt"
	"strings"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"reflect"
)

type Model struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
}

func Upsert(m interface{}, DB *mgo.Database) error {
	return UpsertBy(m, "ID", DB)
}

func UpsertBy(model interface{}, fieldStr string, DB *mgo.Database) error {
	// _, err := models.Collection(new(Room), DB).Upsert(bson.M{"": r.ServiceID}, m)
	// return err

	typ := reflect.TypeOf(model)
	val := reflect.ValueOf(model)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	field, found := typ.FieldByName(fieldStr)

	if !found {
		return fmt.Errorf("Field '%s' not found", fieldStr)
	}

	dbName := strings.Split(field.Tag.Get("bson"), ",")[0]

	// Default it to the lowercase field name if it wasn't provide by the
	// BSON tag
	if dbName == "" {
		dbName = strings.ToLower(field.Name)
	}

	b := bson.M{dbName: val.FieldByName(fieldStr)}
	_, err := Collection(reflect.New(typ).Interface(), DB).Upsert(b, model)

	return err
}

func Collection(model interface{}, DB *mgo.Database) *mgo.Collection {
	typ := reflect.TypeOf(model)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	path := strings.Split(typ.PkgPath(), "models/")
	modelPkg := ""
	if len(path) > 1 {
		modelPkg = strings.Replace(path[1], "/", "_", -1) + "_"
	}

	modelName := strings.ToLower(modelPkg) + strings.ToLower(typ.Name())
	return DB.C(modelName)
}
