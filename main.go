package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	type comment struct {
		Message string `json:"message"`
		Name    string `json:"name"`
	}
	comments := []comment{}

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			_, msgByte, err := conn.ReadMessage()
			if err != nil {
				return
			}

			c := comment{}
			if err := json.Unmarshal(msgByte, &c); err != nil {
				fmt.Println("err json fmt:", err)
				return
			}

			comments = append(comments, c)
		}
	})
	http.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		ticker := time.NewTicker(10 * time.Millisecond)
		for range ticker.C {
			msgJSON, err := json.Marshal(comments)
			if err != nil {
				fmt.Println("err json fmt:", err)
				return
			}

			// Write message back to browser
			if err = conn.WriteMessage(websocket.TextMessage, msgJSON); err != nil {
				return
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":8080", nil)
}
