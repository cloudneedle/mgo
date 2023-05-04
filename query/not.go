package query

import "go.mongodb.org/mongo-driver/bson"

func Not(values ...bson.D) bson.D {
	return bson.D{{Key: "$not", Value: values}}
}
