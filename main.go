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
		database.Test()
	}
	fmt.Println(time.Since(start))
	// auctions.AllPagesAuctions()
}
