package query

import "go.mongodb.org/mongo-driver/bson"

func Or(values ...bson.D) bson.D {
	return bson.D{{Key: "$or", Value: values}}
}
