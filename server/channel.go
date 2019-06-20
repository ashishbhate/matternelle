package main

import "errors"

const channelIdKey = "channelId"

func (p *Plugin) StoreChannelId(channelId string) error {
	return p.API.KVSet(channelIdKey, []byte(channelId))
}

func (p *Plugin) GetChannelId() (string, error) {
	channelId, err := p.API.KVGet(channelIdKey)
	if err != nil {
		return "", err
	}
	if channelId != nil {
		return string(channelId), nil
	}
	return "", errors.New("channel id not found in KVStore")
}
