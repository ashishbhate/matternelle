package main

import (
	"github.com/mattermost/mattermost-server/plugin"
)

func main() {
	p := &Plugin{
		Users:        []*AppUser{},
		Applications: []*App{},
	}
	p.StartWebSocket()
	plugin.ClientMain(p)
}
