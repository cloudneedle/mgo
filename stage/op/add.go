package op

import "go.mongodb.org/mongo-driver/bson"

// Add 将数字加在一起或添加数字和日期。如果参数之一是日期，则$add将其他参数视为毫秒，以添加到日期中。
//
//	{ $add: [ <expression1>, <expression2>, ... ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/add/
func Add(values ...interface{}) bson.D {
	return bson.D{{"$add", values}}
}
