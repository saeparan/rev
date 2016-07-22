package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Email       string        `bson:"Email"`
	Nickname    string        `bson:"Nickname"`
	Password    []byte        `bson:"Password"`
	AccessToken string        `bson:"AccessToken"`
}

func (b *User) Save(s *mgo.Session) error {
	coll := Collection(s, "users")
	err := coll.Insert(b)
	return err
}

func (b *User) Count(s *mgo.Session) (int, error) {
	coll := Collection(s, "users")
	c, err := coll.Count()
	return c, err
}

func (b *User) GetUserByEmail(s *mgo.Session, email string) *User {
	coll := Collection(s, "users")

	u := new(User)
	coll.Find(bson.M{"Email": email}).One(u)
	return u
}
