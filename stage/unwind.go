package stage

import "go.mongodb.org/mongo-driver/bson"

type UnwindField struct {
	Path                       string // 必选，指定要展开的数组字段
	PreserveNullAndEmptyArrays bool   // 可选，指定是否保留空数组，默认为 false
	IncludedArrayIndex         string // 可选，指定一个新字段，用于存储数组元素的索引值
}

// Unwind 展开数组
// https://docs.mongodb.com/manual/reference/operator/aggregation/unwind/
func Unwind(field UnwindField) bson.D {
	var unwind bson.D
	if field.Path != "" {
		unwind = append(unwind, bson.E{Key: "path", Value: field.Path})
	}

	if field.PreserveNullAndEmptyArrays {
		unwind = append(unwind, bson.E{Key: "preserveNullAndEmptyArrays", Value: field.PreserveNullAndEmptyArrays})
	}

	if field.IncludedArrayIndex != "" {
		unwind = append(unwind, bson.E{Key: "includeArrayIndex", Value: field.IncludedArrayIndex})
	}

	return bson.D{{Key: "$unwind", Value: unwind}}
}
