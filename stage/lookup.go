package stage

import "go.mongodb.org/mongo-driver/bson"

type LookupOpt struct {
	From         string
	LocalField   string
	ForeignField string
	As           string
}

// Lookup 联表查询，
//
//	$lookup，https://docs.mongodb.com/manual/reference/operator/aggregation/lookup/
func Lookup(value LookupOpt) bson.D {
	return bson.D{{Key: "$lookup", Value: value}}
}
