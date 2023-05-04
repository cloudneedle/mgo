package query

import "go.mongodb.org/mongo-driver/bson"

func And(values ...bson.D) bson.D {
	return bson.D{{Key: "$and", Value: values}}
}
