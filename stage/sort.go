package stage

import "go.mongodb.org/mongo-driver/bson"

func Sort(values ...bson.D) bson.D {
	var arr = make(bson.D, 0)
	for _, value := range values {
		arr = append(arr, value...)
	}
	return bson.D{{Key: "$sort", Value: arr}}
}
