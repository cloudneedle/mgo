package op

import "go.mongodb.org/mongo-driver/bson"

// Eq 比较两个值是否相等，如果相等则返回 true，否则返回 false。
//
//	{ $eq: [ <expression1>, <expression2> ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/eq/
func Eq(values ...interface{}) bson.D {
	return bson.D{{"$eq", values}}
}
