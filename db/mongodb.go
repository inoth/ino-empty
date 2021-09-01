package db

import (
	"context"
	"ino-empty/config"

	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "github.com/sirupsen/logrus"
)

var (
	mogo *mongo.Database
)

func init() {
	mongoClient, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Instance().MongoDB.Host))
	if err != nil {
		log.Errorf("%v", err)
		os.Exit(1)
	}
	if err = mongoClient.Ping(context.TODO(), nil); err != nil {
		log.Errorf("%v", err)
		os.Exit(1)
	}

	db := mongoClient.Database(config.Instance().MongoDB.DataBase)
	mogo = db
}

func GetDb() *mongo.Database {
	return mogo
}
