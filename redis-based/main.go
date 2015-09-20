package main

import (
	"fmt"
	"log"

	"github.com/timehop/jimmy/redis"
)

func main() {
	db, err := redis.NewPool("redis://localhost:6379", redis.DefaultConfig)
	if err != nil {
		log.Fatal(err.Error())
	}
	foo, err := db.Get("foo")
	fmt.Println(foo, err)
}
