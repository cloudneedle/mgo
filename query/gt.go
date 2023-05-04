package query

import "go.mongodb.org/mongo-driver/bson"

func Gt(value interface{}) bson.D {
	return bson.D{{Key: "$gt", Value: value}}
}

func Gte(value interface{}) bson.D {
	return bson.D{{Key: "$gte", Value: value}}
}
