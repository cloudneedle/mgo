package stage

import "go.mongodb.org/mongo-driver/bson"

// Count 统计
// https://docs.mongodb.com/manual/reference/operator/aggregation/count/
func Count(value string) bson.D {
	return bson.D{{Key: "$count", Value: value}}
}
