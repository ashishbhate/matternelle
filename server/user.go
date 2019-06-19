package main

import (
	"strconv"
)

const nbUsersChatKey = "nbChatUser"

func (p *Plugin) AddChatUser() error {
	return p.updateNbChatUsers(func(value int) int {
		return value + 1
	})
}

func (p *Plugin) RemoveChatUser() error {
	return p.updateNbChatUsers(func(value int) int {
		return value - 1
	})
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
		return err
	}
	return p.API.KVSet(nbUsersChatKey, []byte(strconv.Itoa(transform(nbUsersChat))))
}
