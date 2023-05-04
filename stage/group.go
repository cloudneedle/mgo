package stage

import "go.mongodb.org/mongo-driver/bson"

// Group 分组，_id是必须的，其他字段是可选的
//
//	{
//	  $group: {
//	    _id: <expression>, // 必须
//	    <field1>: { <accumulator1> : <expression1> },
//	    ...
//	}
//	https://docs.mongodb.com/manual/reference/operator/aggregation/group/
func Group(values ...bson.D) bson.D {
	var fields bson.D
	for _, value := range values {
		fields = append(fields, value...)
	}
	return bson.D{{Key: "$group", Value: fields}}
}
