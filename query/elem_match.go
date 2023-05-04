package query

import "go.mongodb.org/mongo-driver/bson"

func ElemMatch(values ...bson.D) bson.D {
	var elemMatch bson.D
	if len(values) == 1 {
		elemMatch = values[0]
	} else {
		elemMatch = And(values...)
	}
	return bson.D{{Key: "$elemMatch", Value: elemMatch}}
}
