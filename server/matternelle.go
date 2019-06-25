package main

import (
	"github.com/pkg/errors"
)

func (p *Plugin) NewAppUser(u *AppUser) error {
	if _, err := p.PostUserMessage(u, "New app user connected"); err != nil {
		return errors.Wrap(err, "can't post msg of new user app to MM")
	}
	p.Users = append(p.Users, u)
	return nil
}

func (p *Plugin) NewAppUserToken(u *AppUser, appUserToken string) error {
	u.Token = appUserToken
	return nil
}

func (p *Plugin) AppUserLeave(u *AppUser) error {
	if _, err := p.PostUserMessage(u, "App user disconnected"); err != nil {
		return errors.Wrap(err, "can't post msg of user app leave to MM")
	}
	var newUsers []*AppUser
	for _, user := range p.Users {
		if u != user {
			newUsers = append(newUsers, user)
		}
	}
	p.Users = newUsers
	return nil
}

func (p *Plugin) NewMessageFromAppUser(user *AppUser, msg string) error {
	postID, err := p.PostUserMessage(user, msg)
	if err != nil {
		return errors.Wrap(err, "can't post msg of user app msg to MM")
	}
	if user.postId == "" {
		user.postId = postID
	}
	return nil
}
