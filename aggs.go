package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Aggs struct {
	pipes []bson.D
	coll  *mongo.Collection
}

func (s *Aggs) Stage(pipes ...bson.D) *Aggs {
	if len(pipes) > 0 {
		for _, v := range pipes {
			if len(v) > 0 {
				s.pipes = append(s.pipes, v)
			}
		}
	}
	return s
}

func (s *Aggs) Find(ctx context.Context, data any) error {
	cur, err := s.coll.Aggregate(ctx, s.pipes)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, data)
	return err
}

func (s *Aggs) FindOne(ctx context.Context, data any) error {
	cur, err := s.coll.Aggregate(ctx, s.pipes)
	if err == mongo.ErrNoDocuments {
		return nil
	}
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	if cur.TryNext(ctx) {
		err = cur.Decode(data)
	}
	return err
}

type PageOption struct {
	PageIndex int64
	PageSize  int64
}

func (s *Aggs) FindPageList(ctx context.Context, opt PageOption, data any) (int64, error) {
	if opt.PageIndex < 1 {
		opt.PageIndex = 1
	}
	if opt.PageSize < 1 {
		opt.PageSize = 10
	}

	count, err := s.Count(ctx)
	if err != nil {
		return 0, err
	}

	var pipes = s.pipes
	if count > 0 {
		s.pipes = append(s.pipes, bson.D{{"$skip", (opt.PageIndex - 1) * opt.PageSize}})
		s.pipes = append(s.pipes, bson.D{{"$limit", opt.PageSize}})
	}

	err = s.Find(ctx, data)
	if err != nil {
		return 0, err
	}
	s.pipes = pipes
	return count, nil
}

func (s *Aggs) Count(ctx context.Context) (int64, error) {
	var pipes = s.pipes
	s.pipes = append(s.pipes, bson.D{{"$count", "count"}})
	var data []map[string]int64
	err := s.Find(ctx, &data)
	if err != nil {
		return 0, err
	}
	s.pipes = pipes
	if data != nil && len(data) > 0 {
		return data[0]["count"], nil
	}
	return 0, nil
}

func (c *Collection) Aggs(pipes ...bson.D) *Aggs {
	var _pipes []bson.D
	if len(pipes) > 0 {
		for _, v := range pipes {
			if len(v) > 0 {
				_pipes = append(_pipes, v)
			}
		}
	}

	return &Aggs{
		coll:  c.coll,
		pipes: _pipes,
	}
}
