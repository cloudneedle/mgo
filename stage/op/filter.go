package op

import "go.mongodb.org/mongo-driver/bson"

type FilterOpt struct {
	Input interface{} //解析为数组的expression。
	As    string      //可选的。 variable的名称，代表input数组的每个单独元素。如果未指定名称，则变量名称默认为this
	Cond  interface{} // expression解析为一个布尔值，该布尔值用于确定是否应在输出数组中包含一个元素。该表达式使用as中指定的变量名称分别引用input数组的每个元素。
}

// Filter 返回一个数组，其中包含输入数组中满足条件的所有元素。
//
//	{ $filter: { input: <expression>, as: <string>, cond: <expression> } }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/filter/
func Filter(value FilterOpt) bson.D {
	return bson.D{{"$filter", value}}
}
