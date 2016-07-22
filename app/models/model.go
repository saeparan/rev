package models

import (
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

var (
    Collections map[string]string
)

// Empty struct to embed in models that will provide application default funcs.
type Model struct {
    Id bson.ObjectId `bson:"_id,omitempty"`
}

func Collection(s *mgo.Session, collection string) *mgo.Collection {
    return s.DB("test").C(collection)
}