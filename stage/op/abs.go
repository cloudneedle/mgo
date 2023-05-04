package op

import "go.mongodb.org/mongo-driver/bson"

// Abs 返回一个数字的绝对值。
//
//	{ $abs: <expression> }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/abs/
func Abs(value interface{}) bson.D {
	return bson.D{{"$abs", value}}
}
