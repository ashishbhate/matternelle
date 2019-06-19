package main

import (
	"strconv"
)

const nbUsersChatKey = "nbUsersChat"

func (p *Plugin) AddUserChat() error {
	return p.updateNbUsersChat(func(value int) int {
		return value + 1
	})
}

func (p *Plugin) RemoveUserChat() error {
	return p.updateNbUsersChat(func(value int) int {
		return value - 1
	})
}

func (p *Plugin) GetNbUsersChat() (int, error) {
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

func (p *Plugin) updateNbUsersChat(transform func(int) int) error {
	nbUsersChat, err := p.GetNbUsersChat()
	if err != nil {
		return err
	}
	return p.API.KVSet(nbUsersChatKey, []byte(strconv.Itoa(transform(nbUsersChat))))
}
