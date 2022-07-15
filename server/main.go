package main

import (
	"github.com/mattermost/mattermost-server/v6/plugin"
)

func main() {
	p := &Plugin{
		Users:        []*AppUser{},
		Applications: []*App{},
	}
	plugin.ClientMain(p)
}
