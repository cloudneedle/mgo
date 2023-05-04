package query

import "go.mongodb.org/mongo-driver/bson"

func Exists(key string, exists bool) bson.D {
	return bson.D{{Key: key, Value: bson.D{{Key: "$exists", Value: exists}}}}
}
