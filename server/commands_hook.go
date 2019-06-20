package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

func (p *Plugin) registerCommand() error {
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          "matternelle",
		DisplayName:      "Matternelle",
		AutoComplete:     true,
		AutoCompleteHint: "[command]",
		AutoCompleteDesc: "Available command : init",
	}); err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	return nil
}

// ExecuteCommand executes a command that has been previously registered via the RegisterCommand
// API.
//
// This demo implementation responds to a /demo_plugin command, allowing the user to enable
// or disable the demo plugin's hooks functionality (but leave the command and webapp enabled).
func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {

	if strings.HasPrefix(args.Command, "/matternelle init") {
		p.StoreChannelId(args.ChannelId)
		p.PostPluginMessage(nil, fmt.Sprintf("Init plugin: %s", args.Command))
		return &model.CommandResponse{}, nil
	} else if strings.HasPrefix(args.Command, "/matternelle on") {
		p.AddChatUser()
		p.PostPluginMessage(nil, fmt.Sprintf("Starting chat: %s", args.Command))
		return &model.CommandResponse{}, nil
	} else if strings.HasPrefix(args.Command, "/matternelle off") {
		p.RemoveChatUser()
		p.PostPluginMessage(nil, fmt.Sprintf("Finish chat: %s", args.Command))
		return &model.CommandResponse{}, nil
	}
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         fmt.Sprintf("Unknown command: %s", args.Command),
	}, nil
}
