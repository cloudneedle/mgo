package op

import "go.mongodb.org/mongo-driver/bson"

type ReduceOpt struct {
	Input        interface{} //解析为数组的expression。
	InitialValue interface{} //初始值，用于存储每个组的累加值。
	In           interface{} //expression解析为一个布尔值，该布尔值用于确定是否应在输出数组中包含一个元素。该表达式使用as中指定的变量名称分别引用input数组的每个元素。
}

// Reduce 通过将每个组的每个元素与累加器值（initialValue）一起传递给指定的表达式来计算每个组的结果。
//
//	{ $reduce: { input: <expression>, initialValue: <expression>, in: <expression> } }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/reduce/
func Reduce(value ReduceOpt) bson.D {
	var reduce bson.D
	if value.Input != nil {
		reduce = append(reduce, bson.E{"input", value.Input})
	}
	if value.InitialValue != nil {
		reduce = append(reduce, bson.E{"initialValue", value.InitialValue})
	}
	if value.In != nil {
		reduce = append(reduce, bson.E{"in", value.In})
	}
	return bson.D{{"$reduce", reduce}}
}
