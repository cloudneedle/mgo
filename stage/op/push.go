package op

import "go.mongodb.org/mongo-driver/bson"

// Push 返回一个* all *值的数组，这些值是通过将表达式应用于一组按键共享相同组的文档中的每个文档而得到的。
//
//	{ $push: { <field1>: <value1>, ... } }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/push/
func Push(value interface{}) bson.D {
	return bson.D{{"$push", value}}
}
