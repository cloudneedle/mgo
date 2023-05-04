package op

import "go.mongodb.org/mongo-driver/bson"

// Floor 返回小于或等于表达式结果的最大整数。
//
//	{ $floor: <number> }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/floor/
func Floor(value interface{}) bson.D {
	return bson.D{{"$floor", value}}
}
