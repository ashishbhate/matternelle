package main

import (
	"github.com/mattermost/mattermost-server/plugin"
)

func main() {
	p := &Plugin{
		Users:        []*AppUser{},
		Applications: []*App{},
	}
	plugin.ClientMain(p)
}
