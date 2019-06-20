package main

import (
	"github.com/mattermost/mattermost-server/plugin"
)

func main() {
	p := &Plugin{
		Users: []*AppUser{},
	}
	p.StartWebSocket()
	plugin.ClientMain(p)
}
