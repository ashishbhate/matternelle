package main

import (
	"github.com/mattermost/mattermost-server/plugin"
)

func main() {
	p := &Plugin{}
	p.StartWebSocket()
	plugin.ClientMain(p)
}
