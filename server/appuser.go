package main

import (
	"github.com/gorilla/websocket"
)

type AppUser struct {
	p      *Plugin
	c      *websocket.Conn
	postId string
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

func (u *AppUser) SendNbChatUser() error {
	nbChatUsers, err := u.p.GetNbChatUsers()
	if err != nil {
		return err
	}
	return u.c.WriteJSON(Command{Command: "nbChatUser", NbChatUser: nbChatUsers})
}

func (u *AppUser) SendMessage(msg string) error {
	return u.c.WriteJSON(Command{Command: "msg", Msg: msg})
}
