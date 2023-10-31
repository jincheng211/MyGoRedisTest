package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:     "mymaster",
		SentinelAddrs:  []string{"sentinel1:26379", "sentinel2:26380", "sentinel3:26381"},
		RouteByLatency: true,
	})

	ctx := context.Background()

	err := client.Set(ctx, "example_key", "example_value", 10*time.Minute).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
	}

	value, err := client.Get(ctx, "example_key").Result()
	if err == redis.Nil {
		fmt.Println("Key not found")
	} else if err != nil {
		fmt.Println("Error getting key:", err)
	} else {
		fmt.Println("example_key:", value)
	}

	client.Close()
}
