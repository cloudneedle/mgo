package update

import "go.mongodb.org/mongo-driver/bson"

func Merge(values ...bson.D) bson.D {
	var d = make(bson.D, 0, len(values))
	for _, value := range values {
		d = append(d, value...)
	}

	return d
}
