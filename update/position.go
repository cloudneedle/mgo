package update

import "go.mongodb.org/mongo-driver/bson"

func Position(value int) bson.D {
	return bson.D{{"$position", value}}
}
