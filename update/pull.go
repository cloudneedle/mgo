package update

import "go.mongodb.org/mongo-driver/bson"

func Pull(values ...bson.D) bson.D {
	var result = make(bson.D, 0)
	for _, value := range values {
		result = append(result, value...)
	}
	return bson.D{{"$pull", result}}
}
