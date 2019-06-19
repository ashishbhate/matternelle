package main

import "github.com/gorilla/websocket"

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
