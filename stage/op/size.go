package op

import "go.mongodb.org/mongo-driver/bson"

// Size 返回数组中的元素总数。
//
//	{ $size: expression }
//	https://docs.mongodb.com/manual/reference/operator/aggregation/size/
func Size(value interface{}) bson.D {
	return bson.D{{"$size", value}}
}
