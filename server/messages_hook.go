package main

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
)

func (p *Plugin) MessageHasBeenPosted(c *plugin.Context, post *model.Post) {
	if post.UserId != p.BotUserID {
		if user := p.getAppUserFromPostId(post.RootId); user != nil {
			if err := user.SendMessage(post.Message); err != nil {
				p.API.LogError(
					"failed to send message to app user",
					"channel_id", post.ChannelId,
					"user_id", post.UserId,
					"error", err.Error(),
				)
			}
		}
	}
}

func (p *Plugin) getAppUserFromPostId(postId string) *AppUser {
	for _, user := range p.Users {
		if postId != "" && user.postId == postId {
			return user
		}
	}
	return nil
}
