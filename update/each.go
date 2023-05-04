package update

import "go.mongodb.org/mongo-driver/bson"

func Each(values ...interface{}) bson.D {
	return bson.D{{"$each", values}}
}
