package update

import "go.mongodb.org/mongo-driver/bson"

func Max(values ...bson.D) bson.D {
	var result = make(bson.D, 0, len(values))
	for _, value := range values {
		result = append(result, value...)
	}

	return bson.D{{"$max", result}}
}
