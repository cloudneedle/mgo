package stage

import (
	"go.mongodb.org/mongo-driver/bson"
)

type ProjectM = bson.M

// Project 字段映射
func Project(value ...bson.D) bson.D {
	var fields bson.D
	for _, v := range value {
		fields = append(fields, v...)
	}
	return bson.D{{Key: "$project", Value: fields}}
}
