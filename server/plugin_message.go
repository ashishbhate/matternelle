package main

import (
	"github.com/mattermost/mattermost-server/model"
)

func (p *Plugin) PostPluginMessage(msg string) error {
	channelId, err := p.GetChannelId()
	if err != nil {
		return err
	}
	_, err = p.API.CreatePost(&model.Post{
		UserId:    p.BotUserID,
		ChannelId: channelId,
		Message:   msg,
	})
	return err
}
