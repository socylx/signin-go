package mongo

import (
	"context"
	"fmt"
	"gsteps-go/global/config"
	"gsteps-go/internal/errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client
var Mongo *mongo.Database

func Init() {
	log.Println("global.mongo.Init Start...")

	dsn := fmt.Sprintf(
		"mongodb://%s:%s@%s:%s/%s",
		config.Mongo.User,
		config.Mongo.Password,
		config.Mongo.Host,
		config.Mongo.Port,
		config.Mongo.Auth,
	)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		log.Fatalf("global.mongo.Init.Connect() Error: %v", errors.Wrap(err, fmt.Sprintf("[DB connection failed] Database name: %s", config.Mysql.Database)))
	}

	err = MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("global.mongo.Init.Ping() Error: %v", err)
	}
	Mongo = MongoClient.Database(config.Mongo.Database)
}

func Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := MongoClient.Disconnect(ctx)
	if err != nil {
		log.Printf("global.mongo.Close Error: %v", err)
	}
}
