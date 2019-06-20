package main

import (
	"github.com/pkg/errors"
	"strconv"
)

const nbUsersChatKey = "nbChatUser"

func (p *Plugin) AddChatUser() error {
	err := p.updateNbChatUsers(func(value int) int {
		return value + 1
	})
	if err != nil {
		return errors.Wrap(err, "can't add chat user")
	}
	return p.sendNbChatUsersToAllAppUsers()
}

func (p *Plugin) RemoveChatUser() error {
	err := p.updateNbChatUsers(func(value int) int {
		return value - 1
	})
	if err != nil {
		return errors.Wrap(err, "can't remove chat user")
	}
	return p.sendNbChatUsersToAllAppUsers()
}

func (p *Plugin) GetNbChatUsers() (int, error) {
	value, err := p.API.KVGet(nbUsersChatKey)
	if err != nil {
		return 0, err
	}
	nbUsersChat := 0
	if value != nil {
		nbUsersChat2, err2 := strconv.Atoi(string(value))
		if err2 != nil {
			return 0, err2
		}
		nbUsersChat = nbUsersChat2
	}
	return nbUsersChat, nil
}

func (p *Plugin) updateNbChatUsers(transform func(int) int) error {
	nbUsersChat, err := p.GetNbChatUsers()
	if err != nil {
		return errors.Wrap(err, "can't get number of chat users")
	}
	return p.API.KVSet(nbUsersChatKey, []byte(strconv.Itoa(transform(nbUsersChat))))
}

func (p *Plugin) sendNbChatUsersToAllAppUsers() error {
	for _, user := range p.Users {
		err := user.SendNbChatUser()
		if err != nil {
			return errors.Wrap(err, "can't send nb of chat users")
		}
	}
	return nil
}
