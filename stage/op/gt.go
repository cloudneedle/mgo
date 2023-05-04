package op

import "go.mongodb.org/mongo-driver/bson"

// Gt 比较两个值，如果第一个值大于第二个值，则返回 true，否则返回 false。
//
//	语法 { $gt: [ <expression1>, <expression2> ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/gt/
func Gt(values ...interface{}) bson.D {
	return bson.D{{"$gt", values}}
}

// Gte 比较两个值，如果第一个值大于或等于第二个值，则返回 true，否则返回 false。
//
//	 语法 { $gte: [ <expression1>, <expression2> ] }
//		https://docs.mongodb.com/manual/reference/operator/aggregation/gte/
func Gte(values ...interface{}) bson.D {
	return bson.D{{"$gte", values}}
}
