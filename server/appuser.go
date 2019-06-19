package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type AppUser struct {
	p *Plugin
	c *websocket.Conn
}

func NewAppUser(p *Plugin, c *websocket.Conn) *AppUser {
	return &AppUser{
		p: p,
		c: c,
	}
}

func (u *AppUser) Leave() {
	u.p.AppUserLeave(u)
}

func (u *AppUser) SendNbChatUser(nb int) error {
	return u.c.WriteJSON([]byte(fmt.Sprintf(`{"nbChatUser":%d}`, nb)))
}

func (u *AppUser) SendMessage(msg string) error {
	return u.c.WriteJSON([]byte(fmt.Sprintf(`{"msg":"%s"}`, msg)))
}
