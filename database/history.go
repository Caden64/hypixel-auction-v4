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
	"reflect"
	"time"
)

func Test() {

	client, err := connect()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	ctx, ctxCancel := newContext()
	// disconnect when done

	defer func() {
		err = disconnect(client, ctx, ctxCancel)
		if err != nil {
			return
		}
	}()

	db, err := databaseConnection(client)

	collExists, err := checkCurrentMonthYearCollExists(db)

	if err != nil {
		fmt.Printf("Error:, %v", err)
		return
	}

	if !collExists {
		err = addCurrentMonthColl(db)
		if err != nil {
			fmt.Printf("Error:, %v", err)
			return
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
			fmt.Println("No data to update")
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
