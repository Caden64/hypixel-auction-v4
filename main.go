package main

import (
	"fmt"
	"hypixel-auction-v4/database"
	"time"
)

func main() {

	start := time.Now()
	database.Test()
	fmt.Println(time.Since(start))
	// auctions.AllPagesAuctions()
}
