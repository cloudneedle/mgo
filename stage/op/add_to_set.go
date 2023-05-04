package op

import "go.mongodb.org/mongo-driver/bson"

// AddToSet 返回所有唯一值的数组，该值是通过将表达式应用于按键共享同一组文档的一组文档中的每个文档而得出的。未指定输出数组中元素的 Sequences。
//
//	{ $addToSet: <expression> }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/addToSet/
func AddToSet(value interface{}) bson.D {
	return bson.D{{"$addToSet", value}}
}
