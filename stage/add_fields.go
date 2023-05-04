package stage

import "go.mongodb.org/mongo-driver/bson"

// AddFields 添加字段,
// 用于在聚合管道中添加新字段,https://docs.mongodb.com/manual/reference/operator/aggregation/addFields/
func AddFields(values ...bson.D) bson.D {
	var fields bson.D
	for _, value := range values {
		fields = append(fields, value...)
	}
	return bson.D{{Key: "$addFields", Value: fields}}
}
