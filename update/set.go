package update

import "go.mongodb.org/mongo-driver/bson"

func Set(values ...bson.D) bson.D {
	var value = make(bson.D, 0, len(values))
	for _, v := range values {
		value = append(value, v...)
	}

	return bson.D{{"$set", value}}
}

func SetObject(obj any) bson.D {
	return bson.D{{"$set", obj}}
}
