package main

import "fmt"

func (p *Plugin) NewAppUser(u *AppUser) error {
	_, err := p.PostPluginMessage(u, "New app user connected")
	return err
}

func (p *Plugin) AppUserLeave(u *AppUser) error {
	_, err := p.PostPluginMessage(u, "App user disconnected")
	return err
}

func (p *Plugin) NewMessageFromAppUser(user *AppUser, msg string) error {
	postId, err := p.PostPluginMessage(user, fmt.Sprintf("New message from app user: %s", msg))
	user.postId = postId
	return err
}
