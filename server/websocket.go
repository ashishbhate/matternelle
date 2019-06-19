package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func StartWebSocket(p *Plugin) {
	go func() {
		http.HandleFunc("/ws", echo)
		log.Fatal(http.ListenAndServe("0.0.0.0:8989", nil))
	}()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
