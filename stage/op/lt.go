package op

import "go.mongodb.org/mongo-driver/bson"

// Lt 比较两个值，如果第一个值小于第二个值，则返回 true，否则返回 false。
//
//	{ $lt: [ <expression1>, <expression2> ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/lt/
func Lt(values ...interface{}) bson.D {
	return bson.D{{"$lt", values}}
}

// Lte 比较两个值，如果第一个值小于或等于第二个值，则返回 true，否则返回 false。
//
//	{ $lte: [ <expression1>, <expression2> ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/lte/
func Lte(values ...interface{}) bson.D {
	return bson.D{{"$lte", values}}
}
