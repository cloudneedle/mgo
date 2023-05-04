package query

import "go.mongodb.org/mongo-driver/bson"

func In(values ...interface{}) bson.D {
	return bson.D{{Key: "$in", Value: values}}
}

func Nin(values ...interface{}) bson.D {
	return bson.D{{Key: "$nin", Value: values}}
}
