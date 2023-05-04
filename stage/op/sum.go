package op

import "go.mongodb.org/mongo-driver/bson"

// Sum 返回所有输入值的和。
//
//	{ $sum: <expression> } or { $sum: [ <expression1>, <expression2>, ... ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/sum/
func Sum(value interface{}) bson.D {
	return bson.D{{"$sum", value}}
}
