package query

import "go.mongodb.org/mongo-driver/bson"

func Concat(values ...interface{}) bson.D {
	return bson.D{{Key: "$concat", Value: values}}
}
