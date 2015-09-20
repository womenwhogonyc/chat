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
	// Configure and set up redis.
	var err error
	Redis, err = redis.NewPool("redis://localhost:6379", redis.DefaultConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()
	// The chat room website
	router.HandleFunc("/chat/{room_id}", HandleChatRoom).Methods("GET")
	// Handle new messages
	router.HandleFunc("/chat/{room_id}/messages", HandleNewMessage).Methods("POST")
	// Get messages
	router.HandleFunc("/chat/{room_id}/messages", HandleGetMessages).Methods("GET")

	http.Handle("/", router)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err.Error())
	}
}

// MY Docs!!!!!
func HandleChatRoom(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)["room_id"]
	w.Write([]byte(fmt.Sprintf("Chat room '%s' is coming soon!", roomID)))
}

func HandleNewMessage(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)["room_id"]
	r.ParseForm()
	message := r.FormValue("message")
	fmt.Println(message)
}

func HandleGetMessages(w http.ResponseWriter, r *http.Request) {
}
