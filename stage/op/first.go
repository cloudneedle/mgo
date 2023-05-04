package op

import "go.mongodb.org/mongo-driver/bson"

// First 返回将表达式应用到按键共享同一组文档的一组文档中的第一个文档所得到的值。仅在文档按定义的 Sequences 排列时才有意义。
//
//	{ $first: <expression> }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/first/
func First(value interface{}) bson.D {
	return bson.D{{"$first", value}}
}
