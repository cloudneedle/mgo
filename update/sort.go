package update

import "go.mongodb.org/mongo-driver/bson"

func Sort(value int) bson.D {
	return bson.D{{"$sort", value}}
}
