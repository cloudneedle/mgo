package mgo

import "go.mongodb.org/mongo-driver/bson"

func KV(key string, value interface{}) bson.D {
	return bson.D{{Key: key, Value: value}}
}

func MergeKV(values ...bson.D) bson.D {
	var fields bson.D
	for _, value := range values {
		fields = append(fields, value...)
	}
	return fields
}
