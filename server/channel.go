package main

import "github.com/pkg/errors"

const channelIdKey = "channelId"

func (p *Plugin) StoreChannelId(channelId string) error {
	return p.API.KVSet(channelIdKey, []byte(channelId))
}

func (p *Plugin) GetChannelId() (string, error) {
	channelId, err := p.API.KVGet(channelIdKey)
	if err != nil {
		return "", errors.Wrap(err, "can't KVGet GetChannelId")
	}
	if channelId != nil {
		return string(channelId), nil
	}
	return "", errors.New("channel id not found in KVStore")
}
