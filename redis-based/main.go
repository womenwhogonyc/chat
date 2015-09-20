package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"

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

type ChatRoom struct {
	RoomID       string
	Messages     []string
	CurrentIndex int
}

// MY Docs!!!!!
func HandleChatRoom(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)["room_id"]

	t, err := template.ParseFiles("chatroom.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	messages, err := Redis.LRange("messages:"+roomID, -100, -1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	chatroom := ChatRoom{
		RoomID:   roomID,
		Messages: messages,
	}

	t.Execute(w, chatroom)
}

func HandleNewMessage(w http.ResponseWriter, r *http.Request) {
	roomID := mux.Vars(r)["room_id"]
	r.ParseForm()
	message := r.FormValue("message")
	if message == "" {
		return
	}
	_, err := Redis.RPush("messages:"+roomID, message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/chat/"+roomID, http.StatusFound)
}

func HandleGetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	roomID := mux.Vars(r)["room_id"]
	messages, err := Redis.LRange("messages:"+roomID, 0, -1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(strings.Join(messages, "<br/>")))
}
