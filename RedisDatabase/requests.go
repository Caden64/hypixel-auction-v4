package RedisDatabse

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"hypixel-auction-v4/HypixelRequests"
	"hypixel-auction-v4/MojangRequests"
)

func Connect() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "password",
		DB:       0,
	})

	return rdb

}

func GetUser(rdb *redis.Client, uuid string) (string, error) {

	ctx := context.Background()

	name, err := rdb.Get(ctx, uuid).Result()

	if err == redis.Nil {
		fmt.Printf("Adding User\n")
		addUser(rdb, ctx, uuid)
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("No need to add User\n")
	}

	return name, nil

}

func addUser(rdb *redis.Client, ctx context.Context, uuid string) {

	client := HypixelRequests.NewClient()

	name, err := MojangRequests.UUIDToUser(uuid, client)

	if err != nil {
		panic(err)
	}

	rdb.Set(ctx, uuid, name, 0)

}

func Disconnect(rdb *redis.Client) error {
	err := rdb.Close()

	return err
}
