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
	"os"
	"reflect"
	"time"
)

func Test() {

	// credential to log in

	//credential := options.Credential{
	//	Username: "root",
	//	Password: "rootpassword",
	//}

	// logged in
	clientOpts := options.Client().ApplyURI(os.Getenv("MONGODB_CONNSTRING")) // .SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)

	// sanity check

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("unable to reach database")
		log.Fatalf("Program failed: %v\n", err)
	}

	// disconnect when done

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Println("Error disconnecting")
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

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	tdata, err := Time(cursor)

	if err != nil {
		panic(err)
	}

	var docs []interface{}

	dbdata, err := Convert(cursor)

	if !reflect.DeepEqual(tdata, time.Time{}) || len(dbdata) == 0 {
		x := auctions.AllPagesAuctions(tdata)

		if len(x.Auctions) == 0 {
			return nil
		}

		for _, i := range x.Auctions {
			docs = append(docs, bson.D{{"auction", i}, {"timestamp", primitive.NewDateTimeFromTime(time.UnixMilli(x.LastUpdated))}})
		}

		fmt.Printf("amount of data to be added %v\n", len(docs))

		_, err = coll.InsertMany(context.TODO(), docs)
		if err != nil {
			log.Fatalf("Error inserting data: %v\n", err)
		}

	}

	cursor, err = coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}

	FinalData, err := Convert(cursor)
	if err != nil {
		log.Fatalf("Error unable to use convert function: %v", err)
	}

	fmt.Printf("final data %v\n", len(FinalData))

	return docs

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
