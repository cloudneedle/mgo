package stage

import "go.mongodb.org/mongo-driver/bson"

func SortByCount(field string) bson.D {
	return bson.D{{Key: "$sortByCount", Value: field}}
}
