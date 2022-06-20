package main

import (
	"fmt"
	"hypixel-auction-v4/HypixelRequests/auctions"
	database "hypixel-auction-v4/MongoDatabase"
	"os"
	"time"
)

func main() {

	time.Sleep(10 * time.Second)

	start := time.Now()
	if os.Getenv("test") != "" {
		if os.Getenv("UseDB") == "true" {
			database.RemoveAll()
			database.Test()

			for {
				time.Sleep(10 * time.Second)
				database.UpdateData()
			}
			// database.Test()
		}
	}

	auctions.AllPagesAuctions(time.Now())

	fmt.Println(time.Since(start))
	// fmt.Println(auctions.URL)
	// auctions.AllPagesAuctions()
}
