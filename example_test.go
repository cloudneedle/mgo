package mgo

import (
	"context"

	"github.com/cloudneedle/mgo/query"
	"github.com/cloudneedle/mgo/stage"

	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	db  *DB
	cli *Client
}

func (s *ExampleTestSuite) SetupSuite() {
	var host = "mongodb://root:123456@10.0.8.3:32193/?directConnection=true"
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	cli, err := NewClient(ctx, host)
	s.Require().NoError(err)
	s.Require().NotNil(cli)
	s.cli = cli
	s.db = cli.NewDB("business")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

type User struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	CreatedAt time.Time `bson:"created_at"`
	Version   int       `bson:"version"`
}

// InsertOne
func (s *ExampleTestSuite) TestInsertOne() {
	var user = User{
		Id:        HexId(),
		Name:      "test",
		Age:       18,
		CreatedAt: time.Now(),
	}
	err := s.db.Table("user").InsertOne(context.Background(), user)
	s.Require().NoError(err)
}

// InsertMany
func (s *ExampleTestSuite) TestInsertMany() {
	var users = make([]any, 0)
	users = append(users, User{
		Id:        HexId(),
		Name:      "test2",
		Age:       19,
		CreatedAt: time.Now(),
	}, User{
		Id:        HexId(),
		Name:      "test3",
		Age:       23,
		CreatedAt: time.Now(),
	})
	err := s.db.Table("user").InsertMany(context.Background(), users)
	s.Require().NoError(err)
}

// Filter FindOne
func (s *ExampleTestSuite) TestFilterFindOne() {
	var user User
	filterByName := query.Eq("name", "test")
	err := s.db.Table("user").Filter(filterByName).FindOne(context.Background(), &user)
	s.Require().NoError(err)
	s.Require().NotNil(user)
	s.Require().Equal("test", user.Name)
}

// Filter FindMany
func (s *ExampleTestSuite) TestFilterFindMany() {
	var users []User
	filterByAge := query.Eq("age", query.Gte(18))
	err := s.db.Table("user").Filter(filterByAge).Find(context.Background(), &users)
	s.Require().NoError(err)
	s.Require().NotNil(users)
	s.Require().Equal(3, len(users))
}

// aggs FindOne
func (s *ExampleTestSuite) TestAggsFindOne() {
	var user User
	match := stage.Match(KV("name", "test")).Stage()
	err := s.db.Table("user").
		Aggs(match).
		FindOne(context.Background(), &user)
	s.Require().NoError(err)
	s.Require().NotNil(user)
	s.Require().Equal("test", user.Name)
}

// aggs 排序
func (s *ExampleTestSuite) TestSort() {
	var users []User
	err := s.db.Table("user").
		Aggs(stage.Sort(KV("age", -1))).Find(context.Background(), &users)
	s.Require().NoError(err)
	s.Require().NotNil(users)
	s.Require().Equal(3, len(users))
	s.T().Log(users)
}

// aggs 分页
func (s *ExampleTestSuite) TestPage() {
	var users []User
	count, err := s.db.Table("user").
		Aggs(stage.Sort(KV("age", -1))).
		FindPageList(context.Background(), PageOption{
			PageIndex: 1,
			PageSize:  10,
		}, &users)
	s.Require().NoError(err)
	s.Require().NotNil(users)
	s.Require().Equal(3, len(users))
	s.T().Log(users)
	s.T().Log(count)
}
