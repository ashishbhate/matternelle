package main

import (
	"github.com/mattermost/mattermost-server/plugin"
)

func main() {
	p := &Plugin{}
	StartWebSocket(p)
	plugin.ClientMain(p)
}
