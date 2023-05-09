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
	opt := bson.D{
		{Key: "from", Value: value.From},
		{Key: "localField", Value: value.LocalField},
		{Key: "foreignField", Value: value.ForeignField},
		{Key: "as", Value: value.As},
	}
	return bson.D{{Key: "$lookup", Value: opt}}
}
