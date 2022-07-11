package services

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongodb() *mongo.Client {
	const uri = "mongodb+srv://doandac:doandac@clusterrestapipractice.ibfxd.mongodb.net/gin-practice?retryWrites=true&w=majority"
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")
	return db
}

func DisconnectMongodb(db *mongo.Client) {
	if err := db.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
