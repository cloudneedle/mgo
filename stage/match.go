package stage

import "go.mongodb.org/mongo-driver/bson"

type match struct {
	condition bson.D
}

// bson.D{{"$match", bson.D{{"name", "张三"}}}},
func Match(condition ...bson.D) *match {
	var m match
	for _, c := range condition {
		m.condition = append(m.condition, c...)
	}
	return &m
}

func (m *match) Match(condition ...bson.D) *match {
	for _, c := range condition {
		m.condition = append(m.condition, c...)
	}
	return m
}

func (m *match) Stage() bson.D {
	if len(m.condition) > 0 {
		return bson.D{{Key: "$match", Value: m.condition}}
	}
	return bson.D{}
}
