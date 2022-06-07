package db

import (
	"context"
	"defaultProject/components/config"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgConnect *MongoConnect

type MongoConnect struct {
	mongo *mongo.Client
}

func (m *MongoConnect) Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 连接池
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Cfg.GetString("Mongo.Host")).SetMaxPoolSize(20))
	if err != nil {
		return err
	}
	m.mongo = client
	mgConnect = m
	return nil
}

type mgo struct {
	database   string
	collection string
}

func NewMgo(database, collection string) *mgo {
	return &mgo{
		database,
		collection,
	}
}

func (m *mgo) GetCollection() (*mongo.Collection, error) {
	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return col, nil
}

func (m *mgo) InsertOne(value interface{}) bool {

	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return false
	}
	_, err = col.InsertOne(context.TODO(), value)
	if err != nil {
		return false
	}
	return true
}

func (m *mgo) InsertMany(values []interface{}) bool {
	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return false
	}
	_, err = col.InsertMany(context.TODO(), values)
	if err != nil {
		return false
	}
	return true
}

func (m *mgo) FindOne(filter interface{}, res interface{}) bool {
	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return false
	}
	err = col.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		return false
	}
	return true
}

func (m *mgo) FindMany(skip, limit int64, filter, sort interface{}) ([]bson.M, bool) {
	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return nil, false
	}
	var findoptions *options.FindOptions
	if limit > 0 {
		findoptions = &options.FindOptions{}
		findoptions.SetSkip(limit * (skip - 1))
		findoptions.SetLimit(limit)
		findoptions.SetSort(sort)
	}
	cur, err := col.Find(context.Background(), filter, findoptions)
	if err != nil {
		return nil, false
	}
	r := make([]bson.M, 0)
	ctx := context.TODO()
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var tmp bson.M
		if err = cur.Decode(&tmp); err != nil {
			logrus.Fatal(err)
		}
		r = append(r, tmp)
	}
	return r, true
}

func (m *mgo) Aggregate(pipeline interface{}) ([]map[string]interface{}, bool) {
	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return nil, false
	}
	opts := options.Aggregate()
	cur, err := col.Aggregate(context.Background(), pipeline, opts)
	if err != nil {
		return nil, false
	}
	ctx := context.TODO()
	r := make([]map[string]interface{}, 0)
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var tmp map[string]interface{}
		if err = cur.Decode(&tmp); err != nil {
			logrus.Fatal(err)
		}
		r = append(r, tmp)
	}
	return r, true
}

func (m *mgo) Count(filter interface{}) int {
	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return 0
	}
	c, err := col.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0
	}
	return int(c)
}

func (m *mgo) UpdateOne(filter, update interface{}) bool {
	col, err := mgConnect.mongo.Database(m.database).Collection(m.collection).Clone()
	if err != nil {
		return false
	}
	_, err = col.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false
	}
	return true
}
