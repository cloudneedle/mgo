package update

import "go.mongodb.org/mongo-driver/bson"

func Rename(from string, to string) bson.D {
	return bson.D{{"$rename", bson.D{{from, to}}}}
}
