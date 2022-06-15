package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func connect() (*mongo.Client, error) {
	credential := options.Credential{
		Username: "root",
		Password: "rootpassword",
	}

	// logged in

	clientOpts := options.Client().ApplyURI("mongodb://root:rootpassword@db").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		return nil, err
	}

	return client, nil

}

func connectionCheck(client *mongo.Client) error {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	return nil

}

func databaseConnection(client *mongo.Client) (*mongo.Database, error) {

	err := connectionCheck(client)

	if err != nil {
		return nil, err
	}

	return client.Database("times_updated"), nil

}

func addCurrentMonthColl(db *mongo.Database) error {
	tso := options.TimeSeries().SetTimeField(time.Now().Format("timestamp"))
	opts := options.CreateCollection().SetTimeSeriesOptions(tso)
	err := db.CreateCollection(context.TODO(), time.Now().Format("January2006"), opts)
	if err != nil {
		return err
	}

	return nil

}

func checkCollExists(db *mongo.Database) (bool, error) {
	names, err := db.ListCollectionNames(context.TODO(), bson.D{})

	if err != nil {
		return false, err
	}

	var collExists bool
	for _, i := range names {
		if i == time.Now().Format("January2006") {
			collExists = true
			break
		}
	}

	return collExists, nil

}

func disconnect(client *mongo.Client) error {

	if err := client.Disconnect(context.TODO()); err != nil {
		return err
	}

	return nil

}

func addManyTimeSeries(dataName string, data []interface{}, timeData time.Time, coll *mongo.Collection, ctx context.Context) error {

	var docs []interface{}

	for _, i := range data {
		docs = append(docs, bson.D{{dataName, i}, {"timestamp", primitive.NewDateTimeFromTime(timeData)}})
	}

	_, err := coll.InsertMany(ctx, docs)

	return err
}

func addOneTimeSeries(dataName string, data interface{}, timeData time.Time, coll *mongo.Collection, ctx context.Context) error {

	_, err := coll.InsertOne(ctx, bson.D{{dataName, data}, {"timestamp", primitive.NewDateTimeFromTime(timeData)}})

	return err
}

func delManyTimeSeries() {

}

func delAllTimeSeries() {

}

func delOneTimeSeries() {

}
