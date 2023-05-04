package op

import "go.mongodb.org/mongo-driver/bson"

// Range 返回一个数组，其元素是生成的数字序列。 $range从指定的起始编号开始，按 Sequences 将起始编号增加指定的步长值，直到(但不包括)endpoints 为止，生成序列。
//
//	{ $range: [ <start>, <end>, <step> ] }
//	start: 起始编号。必须是整数。 end: 结束编号。必须是整数。 step: 步长。必须是整数。如果省略，则默认为1。
//	https://docs.mongodb.com/manual/reference/operator/aggregation/range/
func Range(expr ...interface{}) bson.D {
	return bson.D{{"$range", expr}}
}
