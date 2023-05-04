package query

import "go.mongodb.org/mongo-driver/bson"

func All(values ...bson.D) bson.D {
	return bson.D{{Key: "$all", Value: values}}
}
