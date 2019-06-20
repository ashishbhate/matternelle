package main

import (
	"github.com/mattermost/mattermost-server/model"
	"github.com/pkg/errors"
)

func (p *Plugin) PostPluginMessage(user *AppUser, msg string) (string, error) {
	rootID := ""
	if user != nil {
		rootID = user.postId
	}
	channelID, err := p.GetChannelId()
	if err != nil {
		return "", errors.Wrap(err, "can't get channel id in postPluginMessage")
	}
	if channelID == "" {
		return "", nil
	}
	post, err2 := p.API.CreatePost(&model.Post{
		UserId:    p.BotUserID,
		ChannelId: channelID,
		RootId:    rootID,
		Message:   msg,
	})
	if err2 != nil {
		return "", errors.Wrap(err2, "can't create post in postPluginMessage")
	}
	return post.Id, nil
}
