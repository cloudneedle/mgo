package update

import "go.mongodb.org/mongo-driver/bson"

func Min(values ...bson.D) bson.D {
	var value = make(bson.D, 0, len(values))
	for _, v := range values {
		value = append(value, v...)
	}

	return bson.D{{"$min", value}}
}
