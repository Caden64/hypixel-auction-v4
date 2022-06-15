package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getOneUUID() {

}

func getAllMongoD(coll *mongo.Collection, ctx context.Context) (*mongo.Cursor, error) {
	return coll.Find(ctx, bson.D{})

}

func getAllName() {

}

func getAllDate() {

}

func getAllRarity() {

}

func getAllEnchants() {

}

func getAllReforge() {

}
