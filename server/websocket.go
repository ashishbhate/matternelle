package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type connWs struct {
	c *websocket.Conn
}

func (p *Plugin) StartWebSocket() {
	go func() {
		http.HandleFunc("/ws", p.ws)
		log.Fatal(http.ListenAndServe("0.0.0.0:8989", nil))
	}()
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Command struct {
	Command    string `json:"command"`
	Msg        string `json:"msg,omitempty"`
	NbChatUser int    `json:"nbChatUser,omitempty"`
}

func (p *Plugin) ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		p.API.LogError("Error upgrade conn to web socket", "err", err.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Can't upgrade to web socket", StatusCode: http.StatusInternalServerError})
		return
	}
	defer c.Close()
	appUser := NewAppUser(p, c)
	defer appUser.Leave()
	if err := p.NewAppUser(appUser); err != nil {
		p.API.LogError("Error create new user", "err", err.Error())
		p.writeAPIError(w, &APIErrorResponse{ID: "", Message: "Can't create user", StatusCode: http.StatusInternalServerError})
		return
	}
	for {
		cmd := &Command{}
		err := c.ReadJSON(cmd)
		if err != nil {
			p.API.LogError("websocket read json:", "err", err.Error())
			return
		}
		if cmd.Command == "msg" {
			p.NewMessageFromAppUser(appUser, cmd.Msg)
			appUser.SendMessage("server" + cmd.Msg)
		} else {
			appUser.SendMessage("error : unknown command")
		}
	}
}
