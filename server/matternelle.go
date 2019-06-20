package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func (p *Plugin) NewAppUser(u *AppUser) error {
	if _, err := p.PostPluginMessage(u, "New app user connected"); err != nil {
		return errors.Wrap(err, "can't post msg of new user app to MM")
	}
	return nil
}

func (p *Plugin) AppUserLeave(u *AppUser) error {
	if _, err := p.PostPluginMessage(u, "App user disconnected"); err != nil {
		return errors.Wrap(err, "can't post msg of user app leave to MM")
	}
	return nil
}

func (p *Plugin) NewMessageFromAppUser(user *AppUser, msg string) error {
	postID, err := p.PostPluginMessage(user, fmt.Sprintf("New message from app user: %s", msg))
	if err != nil {
		return errors.Wrap(err, "can't post msg of user app msg to MM")
	}
	if user.postId == "" {
		user.postId = postID
	}
	return nil
}
