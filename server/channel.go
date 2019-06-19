package main

const channelIdKey = "channelId"

func (p *Plugin) StoreChannelId(channelId string) error {
	return p.API.KVSet(channelIdKey, []byte(channelId))
}
