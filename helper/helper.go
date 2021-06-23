package helper

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var clientGlobal *mongo.Client
var ctxGlobal context.Context

func ConnectToMongo() {
	uri := "mongodb+srv://dbUserAdmin:dbPassword@clusternam.e3fz5.mongodb.net/GolangTestDB?authSource=admin&replicaSet=ClusterNam-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ctxGlobal = ctx
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	SetMongoClient(client)
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
}

func GetContext() context.Context {
	return ctxGlobal
}
func GetMongoClient() *mongo.Client {
	return clientGlobal
}
func SetMongoClient(client *mongo.Client) {
	clientGlobal = client
}
