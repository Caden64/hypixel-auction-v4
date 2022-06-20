package database

import (
	"fmt"
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

	if err != nil {
		fmt.Printf("Error:, %v", err)
		return
	}

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
	var docs []interface{}

	client, err := connect()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return docs
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

	if err != nil {
		fmt.Printf("Error:, %v", err)
		return docs
	}
	coll := getCurrentMonthYearColl(db)

	// db.

	cursor, err := getAllMongoD(coll, ctx)
	if err != nil {
		panic(err)
	}

	tdata, err := Time(cursor)

	if err != nil {
		panic(err)
	}

	dbdata, err := Convert(cursor)

	if !reflect.DeepEqual(tdata, time.Time{}) || len(dbdata) == 0 {
		x := auctions.AllPagesAuctions(tdata)

		if len(x.Auctions) == 0 {
			fmt.Println("No data to update")
			return nil
		}

		err = addManyAuctionOneFieldTimeSeries("auction", x.Auctions, time.UnixMilli(x.LastUpdated), coll, ctx)

		if err != nil {
			fmt.Printf("Error:, %v", err)
			return nil
		}

	}

	cursor, err = getAllMongoD(coll, ctx)

	if err != nil {
		fmt.Printf("Error:, %v", err)

		return nil
	}

	FinalData, err := Convert(cursor)
	if err != nil {
		log.Fatalf("Error unable to use convert function: %v", err)
	}

	fmt.Printf("final data %v\n", len(FinalData))

	{
		var test map[string]interface{}

		for cursor.Next(ctx) {

			err = cursor.Decode(&test)

			if err != nil {
				fmt.Printf("ERROR: %v", err)
				log.Fatalf("%v", err)
			}

		}
		fmt.Printf("len of coll: %v\n", len(test))

	}

	getCollStats(db, coll, ctx)

	return docs

}

func RemoveAll() {

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

	if err != nil {
		fmt.Printf("Error:, %v", err)
		return
	}
	coll := getCurrentMonthYearColl(db)

	cursor, err := getAllMongoD(coll, ctx)
	if err != nil {
		panic(err)
	}

	{
		var test map[string]interface{}

		for cursor.Next(ctx) {

			err = cursor.Decode(&test)

			if err != nil {
				fmt.Printf("ERROR: %v", err)
				return
			}

		}
		fmt.Printf("len of coll: %v\n", len(test))

	}

	err = delAllTimeSeries(coll, ctx)
	if err != nil {
		fmt.Printf("Error:, %v", err)
		return
	}

	{
		var test map[string]interface{}

		for cursor.Next(ctx) {

			err = cursor.Decode(&test)

			if err != nil {
				fmt.Printf("ERROR: %v", err)
				return
			}

		}
		fmt.Printf("len of coll: %v\n", len(test))

	}

	err = dropTimeSeries(coll, ctx)

	if err != nil {
		fmt.Printf("Error:, %v", err)
		return
	}

}
