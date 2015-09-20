package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timehop/jimmy/redis"
)

var Redis redis.Pool

func main() {
	var err error
	Redis, err = redis.NewPool("redis://localhost:6379", redis.DefaultConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/chat/{room_id}", HandleChatRoom)

	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}

func HandleChatRoom(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)["room_id"]
	w.Write([]byte(fmt.Sprintf("Chat room %s is coming Soon!", roomID)))
}
