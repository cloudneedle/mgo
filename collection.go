package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICollection interface {
	TableName() string
}

type Collection struct {
	coll   *mongo.Collection
	filter *Filter
}

func (c *Collection) Filter(ops ...bson.D) *Filter {
	c.filter = &Filter{coll: c.coll, pipes: bson.D{}}

	for _, op := range ops {
		if op != nil {
			for _, e := range op {
				c.filter.pipes = append(c.filter.pipes, e)
			}
		}
	}

	return c.filter
}

func (c *Collection) InsertOne(ctx context.Context, doc any) error {
	_, err := c.coll.InsertOne(ctx, doc)
	return err
}

func (c *Collection) InsertMany(ctx context.Context, docs []any) error {
	_, err := c.coll.InsertMany(ctx, docs)
	return err
}
