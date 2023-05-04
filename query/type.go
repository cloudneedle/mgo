package query

import "go.mongodb.org/mongo-driver/bson"

func Type(t interface{}) bson.D {
	return bson.D{{Key: "$type", Value: t}}
}
