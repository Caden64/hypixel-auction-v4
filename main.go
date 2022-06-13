package main

import (
	"fmt"
	"hypixel-auction-v4/database"
	"os"
	"time"
)

func main() {

	start := time.Now()
	if os.Getenv("UseDB") == "true" {
		database.RemoveAll()
		database.Test()

		for {
			time.Sleep(20 * time.Second)
			database.UpdateData()
		}
		// database.Test()
	}
	fmt.Println(time.Since(start))
	// fmt.Println(auctions.URL)
	// auctions.AllPagesAuctions()
}
