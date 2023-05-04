package mgo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DB struct {
	db *mongo.Database
}

func (d *DB) Table(name string) *Collection {
	return &Collection{coll: d.db.Collection(name)}
}

func (d *DB) Model(coll ICollection) *Collection {
	return &Collection{coll: d.db.Collection(coll.TableName())}
}

func HexId() string {
	return primitive.NewObjectID().Hex()
}

func ObjectId() primitive.ObjectID {
	return primitive.NewObjectID()
}
