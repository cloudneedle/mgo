package op

import "go.mongodb.org/mongo-driver/bson"

// In 返回一个布尔值，该值指示表达式的结果是否在指定的数组中。
//
//	{ $in: [ <expression>, <arrayExpression> ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/in/
func In(values ...interface{}) bson.D {
	return bson.D{{"$in", values}}
}
