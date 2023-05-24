package query

import "go.mongodb.org/mongo-driver/bson"

type RegexOpt struct {
	Field   string // 字段
	Pattern string // 正则表达式
	Options string // i: 不区分大小写; m: 多行匹配; x: 忽略空白字符; s: 单行匹配;
}

// { <field>: { $regex: /pattern/, $options: '<options>' } }
func Regex(opt RegexOpt) bson.D {
	var value = bson.D{}
	value = append(value, bson.E{Key: "$regex", Value: opt.Pattern})
	if opt.Options != "" {
		value = append(value, bson.E{Key: "$options", Value: opt.Options})
	}
	return bson.D{{opt.Field, value}}
}
