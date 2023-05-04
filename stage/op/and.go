package op

import "go.mongodb.org/mongo-driver/bson"

// And 计算一个或多个表达式，如果所有表达式都是true或如果没有参数表达式则被唤醒，则返回true。否则，$and返回false。
//
//	{ $and: [ { <expression1> }, { <expression2> }, ... , { <expressionN> } ] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/and/
func And(values ...interface{}) bson.D {
	return bson.D{{"$and", values}}
}
