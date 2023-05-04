package op

import "go.mongodb.org/mongo-driver/bson"

// AllElementsTrue 将数组评估为集合，如果数组中的* no *元素为false，则返回true。否则，返回false。空数组将返回true
//
//	{ $allElementsTrue: [<expression>] }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/allElementsTrue/
func AllElementsTrue(values ...interface{}) bson.D {
	return bson.D{{"$allElementsTrue", values}}
}
