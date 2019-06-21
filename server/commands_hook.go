package main

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/mattermost/mattermost-server/model"
	"github.com/mattermost/mattermost-server/plugin"
)

const CommandMatternelle = "/matternelle"
const CommandInit = CommandMatternelle + " init"
const CommandRemove = CommandMatternelle + " remove"
const CommandList = CommandMatternelle + " list"
const CommandOn = CommandMatternelle + " on"
const CommandOff = CommandMatternelle + " off"

func (p *Plugin) registerCommand() error {
	if err := p.API.RegisterCommand(&model.Command{
		Trigger:          "matternelle",
		DisplayName:      "Matternelle",
		AutoComplete:     true,
		AutoCompleteHint: "[command]",
		AutoCompleteDesc: "Available commands: init, on, off",
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

	if strings.HasPrefix(args.Command, CommandInit) {
		appName := strings.TrimSpace(strings.SplitAfter(args.Command, CommandInit)[1])

		if appName == "" {
			return &model.CommandResponse{
				ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
				Text:         "Init command must be followed by the app name",
			}, nil
		}
		app, _ := p.initialize(appName, args.ChannelId)
		p.PostPluginMessage(args.ChannelId, fmt.Sprintf(
			"Init plugin for application %s. Copy and paste the token %s in your web component",
			appName, app.Token))
		return &model.CommandResponse{}, nil
	} else if strings.HasPrefix(args.Command, CommandRemove) {
		appName := strings.TrimSpace(strings.SplitAfter(args.Command, CommandRemove)[1])

		if appName == "" {
			return &model.CommandResponse{
				ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
				Text:         "Remove command must be followed by the app name",
			}, nil
		}
		p.remove(appName)
		p.PostPluginMessage(args.ChannelId, fmt.Sprintf("Application %s removed successfully", appName))
		return &model.CommandResponse{}, nil
	} else if strings.HasPrefix(args.Command, CommandList) {
		msg := []string{"List of applications:"}
		for _, app := range p.Applications {
			msg = append(msg, fmt.Sprintf("%s: %s", app.Name, app.Token))
		}
		p.PostPluginMessage(args.ChannelId, strings.Join(msg, "\n* "))
		return &model.CommandResponse{}, nil
	} else if strings.HasPrefix(args.Command, CommandOn) {
		p.AddChatUser()
		p.PostPluginMessage(args.ChannelId, fmt.Sprintf("Starting chat: %s", args.Command))
		return &model.CommandResponse{}, nil
	} else if strings.HasPrefix(args.Command, CommandOff) {
		p.RemoveChatUser()
		p.PostPluginMessage(args.ChannelId, fmt.Sprintf("Finish chat: %s", args.Command))
		return &model.CommandResponse{}, nil
	}
	return &model.CommandResponse{
		ResponseType: model.COMMAND_RESPONSE_TYPE_EPHEMERAL,
		Text:         fmt.Sprintf("Unknown command: %s", args.Command),
	}, nil
}
