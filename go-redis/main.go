package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	address := os.Getenv("ADDR")
	password := os.Getenv("PASSWORD")
	database, _ := strconv.Atoi(os.Getenv("DATABASE"))

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: database,
	})
	defer rdb.Close()

	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal("Redis connection was refused")
	}
	fmt.Println(status)

	// adding data to the cache
	addData(rdb, "FOO1", "BAR1", 0)
	addData(rdb, "INT", 372, 0)
	addData(rdb, "EXPIRING", "suraj yadav", 5*time.Minute)
	rdb.HSet(ctx, "STRUCT", Person{"John", 25})

	// reading data from the cache
	readData(rdb, "FOO1")
	readData(rdb, "INT")
	readData(rdb, "EXPIRING")
	readStruct(rdb, "STRUCT")

	// update data in cache
	addData(rdb, "FOO1", "BAR2", 0)
	readData(rdb, "FOO1")

	// delete data from cache
	deleteData(rdb, "FOO1")
	readData(rdb, "FOO1")
}

type Person struct {
	NAME string `redis:"name"`
	Age int `redis:"age"`
}

func addData(rdb *redis.Client, key string, value any, duration time.Duration) {
	ctx := context.Background()
	_, err := rdb.Set(ctx, key, value, duration).Result()
	if err != nil {
		fmt.Printf("Failed to add Key(%s)-Value pair\n", key)
		return
	}
}


func readData(rdb *redis.Client, key string) {
	ctx := context.Background()

	result, err := rdb.Get(ctx, key).Result()
	if err != nil {
		fmt.Printf("Key %s not found in Redis Cache\n", key)
		return
	}
	fmt.Printf("%s: %v, %T\n", key, result, result)
}

func readStruct(rdb *redis.Client, key string) {
	ctx := context.Background()

	var person Person
	err := rdb.HGetAll(ctx, key).Scan(&person)
	if err != nil {
		fmt.Printf("%s not found in Redis cache\n", key)
		return
	}
	fmt.Printf("%s: %v\n", key, person)
}

func deleteData(rdb *redis.Client, key string) {
	ctx := context.Background()

	rdb.Del(ctx, key)
}