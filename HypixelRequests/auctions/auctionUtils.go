package auctions

import "sync"

// AuctionData to turn json from requests into usable data
type AuctionData struct {
	Success       bool      `json:"success" bson:"success"`
	Page          int       `json:"page" bson:"page"`
	TotalPages    int       `json:"totalPages" bson:"totalPages"`
	TotalAuctions int       `json:"totalAuctions" bson:"totalAuctions"`
	LastUpdated   int64     `json:"lastUpdated" bson:"lastUpdated"`
	Auctions      []Auction `json:"auctions" bson:"auctions"`
}

// Auction has data some will not be filled by the request because
type Auction struct {
	Uuid             string   `json:"uuid" bson:"uuid"`
	Auctioneer       string   `json:"auctioneer" bson:"auctioneer"`
	ProfileId        string   `json:"profile_id" bson:"profileId"`
	Coop             []string `json:"coop" bson:"coop"`
	CoopUser         []string `json:"coopUser" bson:"coopUser"`
	Start            int64    `json:"start" bson:"start"`
	End              int64    `json:"end" bson:"end"`
	ItemName         string   `json:"item_name" bson:"itemName"`
	ItemLore         string   `json:"item_lore" bson:"itemLore"`
	Extra            string   `json:"extra" bson:"extra"`
	Category         string   `json:"category" bson:"category"`
	Tier             string   `json:"tier" bson:"tier"`
	StartingBid      int32    `json:"starting_bid" bson:"startingBid"`
	Claimed          bool     `json:"claimed" bson:"claimed"`
	HighestBidAmount int32    `json:"highest_bid_amount" bson:"highestBidAmount"`
	Bin              bool     `json:"bin,omitempty" bson:"bin,omitempty"`
	LowestPrice      int32    `bson:"lowestPrice"`
	HighestPrice     int32    `bson:"highestPrice"`
	Reforge          string   `bson:"reforge"`
	Recombobulated   bool     `bson:"recombobulated"`
	Dungeoned        bool     `bson:"dungeoned"`
	DungeonedLvl     int32    `json:"dungeoned_lvl" bson:"dungeonedLvl"`
	LimitedUsage     bool     `json:"limited_usage" bson:"limitedUsage"`
}

type AllAuctionData struct {
	Mutex       sync.Mutex
	Auctions    []Auction
	LastUpdated int64
	Pages       int
}
