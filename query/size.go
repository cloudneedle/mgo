package query

import "go.mongodb.org/mongo-driver/bson"

func Size(size int) bson.D {
	return bson.D{{Key: "$size", Value: size}}
}
