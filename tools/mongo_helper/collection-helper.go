package mongo_helper

import (
	"gopkg.in/mgo.v2"
	"time"
)

func GetMongoCollection(connString, db, collection string) (*mgo.Collection, error) {
	sess, err := mgo.Dial(connString)
	if err != nil {
		return nil, err
	}
	sess.SetSocketTimeout(time.Hour)
	coll := sess.DB(db).C(collection)
	return coll, nil
}
