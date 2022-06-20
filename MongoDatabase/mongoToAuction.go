package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hypixel-auction-v4/HypixelRequests/auctions"
	"reflect"
	"time"
)

func Convert(cursor *mongo.Cursor) ([]auctions.Auction, error) {
	var test map[string]interface{}

	auction := auctions.Auction{}
	var totalAuction []auctions.Auction

	for cursor.Next(context.TODO()) {

		err := cursor.Decode(&test)

		if err != nil {
			return nil, err
		}

		for _, i := range test {

			if reflect.TypeOf(i).Kind() == reflect.Map {

				for m, k := range i.(map[string]interface{}) {
					if m == "uuid" {
						if !reflect.DeepEqual(auction, auctions.Auction{}) {
							totalAuction = append(totalAuction, auction)
						}
						auction.Uuid = k.(string)
					} else if m == "auctioneer" {
						auction.Auctioneer = k.(string)
					} else if m == "profileId" {
						auction.ProfileId = k.(string)
					} else if m == "coop" {
						if pa, ok := k.(primitive.A); ok {
							value := []interface{}(pa)
							var coop []string
							for _, t := range value {
								coop = append(coop, t.(string))
							}
							auction.Coop = coop
						}
					} else if m == "start" {
						auction.Start = k.(int64)
					} else if m == "end" {
						auction.End = k.(int64)
					} else if m == "itemName" {
						auction.ItemName = k.(string)
					} else if m == "itemLore" {
						auction.ItemLore = k.(string)
					} else if m == "extra" {
						auction.Extra = k.(string)
					} else if m == "category" {
						auction.Category = k.(string)
					} else if m == "tier" {
						auction.Tier = k.(string)
					} else if m == "startingBid" {
						auction.StartingBid = k.(int32)
					} else if m == "claimed" {
						auction.Claimed = k.(bool)
					} else if m == "highestBidAmount" {
						auction.HighestBidAmount = k.(int32)
					} else if m == "bin" {
						auction.Bin = k.(bool)
					} else if m == "lowestPrice" {
						auction.LowestPrice = k.(int32)
					} else if m == "highestPrice" {
						auction.HighestPrice = k.(int32)
					} else if m == "reforge" {
						auction.Reforge = k.(string)
					} else if m == "recombobulated" {
						auction.Recombobulated = k.(bool)
					} else if m == "dungeoned" {
						auction.Dungeoned = k.(bool)
					} else if m == "dungeonedLvl" {
						auction.DungeonedLvl = k.(int32)
					} else if m == "limitedUsage" {
						auction.LimitedUsage = k.(bool)
					}

				}
			}
		}
	}

	return totalAuction, nil

}

func Time(cursor *mongo.Cursor) (time.Time, error) {

	var test map[string]interface{}

	for cursor.Next(context.TODO()) {

		err := cursor.Decode(&test)

		if err != nil {
			return time.Time{}, err
		}

	}

	for j, i := range test {
		if j == "timestamp" {
			return i.(primitive.DateTime).Time(), nil
		}
	}

	return time.Time{}, nil

}
