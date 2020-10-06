package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gocs/chatwebsocket/youtube"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return ws, err
	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)

		for t := range ticker.C {
			fmt.Printf("Updating stats: %+v\n", t)

			items, err := youtube.GetSubscribers()
			if err != nil {
				fmt.Println(err)
			}

			jsonStr, err := json.Marshal(items)
			if err != nil {
				fmt.Println(err)
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(jsonStr)); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
