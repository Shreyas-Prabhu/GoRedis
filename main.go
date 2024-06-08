package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func main() {
	fmt.Println("Redis go")
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	products := Product{
		Id:   os.Args[1],
		Name: "Shreyas",
		Type: "Go",
	}

	byteData, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}
	// err = client.Set(products.Id, byteData, 50*time.Second).Err()
	err = client.Set(products.Id, byteData, 0).Err()
	if err != nil {
		panic(err)
	}

	res, err := client.Get(products.Id).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

	num, err := client.Del("1").Result()
	fmt.Println(num, err)

}
