package main

import (
	"github.com/mattermost/mattermost-server/model"
)

func (p *Plugin) PostPluginMessage(user *AppUser, msg string) (string, error) {
	var rootId string
	if user != nil {
		rootId = user.postId
	}
	channelId, err := p.GetChannelId()
	if err != nil {
		return "", err
	}
	post, err := p.API.CreatePost(&model.Post{
		UserId:    p.BotUserID,
		ChannelId: channelId,
		RootId:    rootId,
		Message:   msg,
	})
	return post.Id, err
}
