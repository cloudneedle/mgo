package stage

import "go.mongodb.org/mongo-driver/bson"

// Skip 跳过
// https://docs.mongodb.com/manual/reference/operator/aggregation/skip/
func Skip(n int) bson.D {
	return bson.D{{Key: "$skip", Value: n}}
}
