package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"
)

func (p *Plugin) PostPluginMessage(channelId string, msg string) (string, error) {
	post, err := p.API.CreatePost(&model.Post{
		UserId:    p.BotUserID,
		ChannelId: channelId,
		Message:   msg,
	})
	if err != nil {
		return "", errors.Wrap(err, "can't create post in postPluginMessage")
	}
	return post.Id, nil
}

func (p *Plugin) PostUserMessage(user *AppUser, msg string) (string, error) {
	ChannelId := ""
	for _, app := range p.Applications {
		if user.Token == app.Token {
			ChannelId = app.ChannelId
		}
	}
	if ChannelId == "" {
		return "", fmt.Errorf("No app found for token %s", user.Token)
	}
	post, err2 := p.API.CreatePost(&model.Post{
		UserId:    p.BotUserID,
		ChannelId: ChannelId,
		RootId:    user.postId,
		Message:   msg,
	})
	if err2 != nil {
		return "", errors.Wrap(err2, "can't create post in postPluginMessage")
	}
	return post.Id, nil
}
