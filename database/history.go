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

const uri = "mongodb://root:exsample@localhost"

func Test() {

	// credential to log in

	credential := options.Credential{
		Username: "root",
		Password: "example",
	}

	// logged in

	clientOpts := options.Client().ApplyURI(uri).SetAuth(credential)
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
		fmt.Println(i)
		if i == time.Now().Format("January2006") {
			collExsits = true
			break
		}
	}

	if !collExsits {
		tso := options.TimeSeries().SetTimeField(time.Now().Format("03:04:05 PM 02 January 2006"))
		opts := options.CreateCollection().SetTimeSeriesOptions(tso)

		err = db.CreateCollection(context.TODO(), "June2022", opts)
		if err != nil {
			panic(err)

		}

	}

	coll := db.Collection("June2022")

	_, err = coll.DeleteMany(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("error deleting all data")
	}

	x := auctions.AllPagesAuctions()

	var docs []interface{}

	for _, i := range x.Auctions {
		docs = append(docs, bson.D{{"auction", i}, {"time", primitive.NewDateTimeFromTime(time.UnixMilli(x.LastUpdated))}})
	}

	fmt.Println(len(docs))

	// docs = append(docs, primitive.NewDateTimeFromTime(time.UnixMilli(data.LastUpdated)))
	_, err = coll.InsertMany(context.TODO(), docs)
	if err != nil {
		log.Fatalf("Error: %v from request", err)
	}

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
