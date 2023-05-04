package mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Filter struct {
	pipes                    bson.D
	coll                     *mongo.Collection
	findOptions              []*options.FindOptions
	findOneOptions           []*options.FindOneOptions
	findOneAndDeleteOptions  []*options.FindOneAndDeleteOptions
	findOneAndReplaceOptions []*options.FindOneAndReplaceOptions
	findUpdateOptions        []*options.FindOneAndUpdateOptions
	deleteOptions            []*options.DeleteOptions
	updateOptions            []*options.UpdateOptions
	replaceOptions           []*options.ReplaceOptions
}

func (f *Filter) Filter(op bson.D) *Filter {
	if op != nil {
		for _, e := range op {
			f.pipes = append(f.pipes, e)
		}
	}

	return f
}

func (f *Filter) Find(ctx context.Context, data any) error {
	cur, err := f.coll.Find(ctx, f.pipes, f.findOptions...)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, data)
	return err
}

func (f *Filter) FindOne(ctx context.Context, data any) error {
	err := f.coll.FindOne(ctx, f.pipes, f.findOneOptions...).Decode(data)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func (f *Filter) DeleteOne(ctx context.Context) error {
	_, err := f.coll.DeleteOne(ctx, f.pipes, f.deleteOptions...)
	return err
}

func (f *Filter) DeleteMany(ctx context.Context) error {
	_, err := f.coll.DeleteMany(ctx, f.pipes, f.deleteOptions...)
	return err
}

func (f *Filter) UpdateOne(ctx context.Context, data any) error {
	_, err := f.coll.UpdateOne(ctx, f.pipes, data, f.updateOptions...)
	return err
}

func (f *Filter) UpdateMany(ctx context.Context, data any) error {
	_, err := f.coll.UpdateMany(ctx, f.pipes, data, f.updateOptions...)
	return err
}

func (f *Filter) ReplaceOne(ctx context.Context, data any) error {
	_, err := f.coll.ReplaceOne(ctx, f.pipes, data, f.replaceOptions...)
	return err
}

func (f *Filter) FindOneAndDelete(ctx context.Context, data any) error {
	err := f.coll.FindOneAndDelete(ctx, f.pipes, f.findOneAndDeleteOptions...).Decode(data)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func (f *Filter) FindOneAndReplace(ctx context.Context, update, data any) error {
	err := f.coll.FindOneAndReplace(ctx, f.pipes, update, f.findOneAndReplaceOptions...).Decode(data)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func (f *Filter) FindOneAndUpdate(ctx context.Context, update, data any) error {
	err := f.coll.FindOneAndUpdate(ctx, f.pipes, update, f.findUpdateOptions...).Decode(data)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

func (f *Filter) CountDocuments(ctx context.Context) (int64, error) {
	return f.coll.CountDocuments(ctx, f.pipes)
}

func (f *Filter) FindOptions(opts ...*options.FindOptions) *Filter {
	f.findOptions = append(f.findOptions, opts...)
	return f
}

func (f *Filter) FindOneOptions(opts ...*options.FindOneOptions) *Filter {
	f.findOneOptions = append(f.findOneOptions, opts...)
	return f
}

func (f *Filter) FindOneAndDeleteOptions(opts ...*options.FindOneAndDeleteOptions) *Filter {
	f.findOneAndDeleteOptions = append(f.findOneAndDeleteOptions, opts...)
	return f
}

func (f *Filter) FindOneAndReplaceOptions(opts ...*options.FindOneAndReplaceOptions) *Filter {
	f.findOneAndReplaceOptions = append(f.findOneAndReplaceOptions, opts...)
	return f
}

func (f *Filter) FindOneAndUpdateOptions(opts ...*options.FindOneAndUpdateOptions) *Filter {
	f.findUpdateOptions = append(f.findUpdateOptions, opts...)
	return f
}

func (f *Filter) DeleteOptions(opts ...*options.DeleteOptions) *Filter {
	f.deleteOptions = append(f.deleteOptions, opts...)
	return f
}

func (f *Filter) UpdateOptions(opts ...*options.UpdateOptions) *Filter {
	f.updateOptions = append(f.updateOptions, opts...)
	return f
}

func (f *Filter) ReplaceOptions(opts ...*options.ReplaceOptions) *Filter {
	f.replaceOptions = append(f.replaceOptions, opts...)
	return f
}
