package query

import "go.mongodb.org/mongo-driver/bson"

func Eq(key string, value interface{}) bson.D {
	return bson.D{{Key: key, Value: value}}
}

func Ne(key string, value interface{}) bson.D {
	return bson.D{{Key: key, Value: bson.M{"$ne": value}}}
}
