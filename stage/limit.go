package stage

import "go.mongodb.org/mongo-driver/bson"

func Limit(n int) bson.D {
	return bson.D{{Key: "$limit", Value: n}}
}
