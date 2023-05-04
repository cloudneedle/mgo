package mgo

import "go.mongodb.org/mongo-driver/bson"

func KV(key string, value interface{}) bson.D {
	return bson.D{{Key: key, Value: value}}
}
