package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hypixel-auction-v4/HypixelRequests/auctions"
	"log"
	"time"
)

func Test() {

	// credential to log in

	credential := options.Credential{
		Username: "root",
		Password: "rootpassword",
	}

	// logged in

	clientOpts := options.Client().ApplyURI("mongodb://root:rootpassword@db").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	// sanity check

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("unable to reach database")
		log.Fatalf("Program failed: %v\n", err)
	}

	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	// disconnect when done

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Println("error")
			panic(err)
		}
	}()

	db := client.Database("times_updated")

	names, err := db.ListCollectionNames(context.TODO(), bson.D{})
	var collExsits bool
	for _, i := range names {
		if i == time.Now().Format("January2006") {
			collExsits = true
			break
		}
	}

	if !collExsits {
		tso := options.TimeSeries().SetTimeField(time.Now().Format("timestamp"))
		opts := options.CreateCollection().SetTimeSeriesOptions(tso)

		err = db.CreateCollection(context.TODO(), time.Now().Format("January2006"), opts)
		if err != nil {
			panic(err)

		}

	}

	coll := db.Collection(time.Now().Format("January2006"))

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	Finaldata, err := Convert(cursor)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(len(Finaldata))

}

func UpdateData() []interface{} {

	credential := options.Credential{
		Username: "root",
		Password: "rootpassword",
	}

	// logged in

	clientOpts := options.Client().ApplyURI("mongodb://root:rootpassword@db").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	db := client.Database("times_updated")

	coll := db.Collection(time.Now().Format("January2006"))

	_, err = coll.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("error deleting all data")
	}

	x := auctions.AllPagesAuctions()

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	TestData, err := Convert(cursor)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// if len(TestData) != 0 && check(TestData)

	var docs []interface{}

	for _, i := range x.Auctions {
		docs = append(docs, bson.D{{"auction", i}, {"timestamp", primitive.NewDateTimeFromTime(time.UnixMilli(x.LastUpdated))}})
	}

	fmt.Println(len(docs))

	_, err = coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatalf("Error: %v from request", err)
	}

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	FinalData, err := Convert(cursor)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println(len(FinalData))

	return docs

}

func check(first, second time.Time) bool {

	if first != second {
		return false
	}

	return true
}

func RemoveAll() {

	credential := options.Credential{
		Username: "root",
		Password: "rootpassword",
	}

	// logged in

	clientOpts := options.Client().ApplyURI("mongodb://root:rootpassword@db").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	db := client.Database("times_updated")

	coll := db.Collection(time.Now().Format("January2006"))

	_, err = coll.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("error deleting all data")
	}

	err = coll.Drop(context.TODO())
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
}
