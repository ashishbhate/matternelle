package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

type connWs struct {
	c *websocket.Conn
}

func (p *Plugin) StartWebSocket() error {
	port, err := strconv.ParseInt(p.configuration.WebSocketPort, 10, 0)
	if err != nil {
		return err
	}
	go func() {
		http.HandleFunc("/ws", p.ws)

		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
	}()

	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Command struct {
	Command      string `json:"command"`
	Msg          string `json:"msg,omitempty"`
	NbChatUser   int    `json:"nbChatUser,omitempty"`
	AppUserToken string `json:"appUserToken,omitempty"`
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
	appUser.SendNbChatUser()
	defer appUser.Leave()
	err = p.NewAppUser(appUser)
	if err != nil {
		p.API.LogError("Error create new user", "err", err.Error())
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
			if appUser.Token != "" {
				if err := p.NewMessageFromAppUser(appUser, cmd.Msg); err != nil {
					p.API.LogError("err new message from app user", "err", err.Error())
				}
			} else {
				p.API.LogError("can't deliver msg because no token", "msg", cmd.Msg)
			}
		} else if cmd.Command == "tokenApp" {
			if err := p.NewAppUserToken(appUser, cmd.AppUserToken); err != nil {
				p.API.LogError("err new token app user", "err", err.Error())
			}
		} else {
			appUser.SendMessage("error : unknown command")
		}
	}
}
