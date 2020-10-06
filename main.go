package main

import (
	"fmt"
	"log"
	"net/http"

	socket "github.com/gocs/chatwebsocket/websocket"
)

func main() {

	fmt.Println("running...")
	setupRoutes()
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/stats", stats)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "home page")
}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := socket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go socket.Writer(ws)
}
