package query

import "go.mongodb.org/mongo-driver/bson"

func Nor(values ...bson.D) bson.D {
	return bson.D{{Key: "$nor", Value: values}}
}
