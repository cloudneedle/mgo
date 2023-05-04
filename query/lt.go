package query

import "go.mongodb.org/mongo-driver/bson"

func Lt(value interface{}) bson.D {
	return bson.D{{Key: "$lt", Value: value}}
}

func Lte(value interface{}) bson.D {
	return bson.D{{Key: "$lte", Value: value}}
}
