module github.com/ashishbhate.com/matternelle

go 1.16

require (
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/mattermost/mattermost-plugin-api v0.0.27 // indirect
	github.com/mattermost/mattermost-server/v6 v6.7.2
	github.com/pkg/errors v0.9.1
)

// Workaround for https://github.com/golang/go/issues/30831 and fallout.
replace github.com/golang/lint => github.com/golang/lint v0.0.0-20181217174547-8f45f776aaf1

// To dev https://github.com/mattermost/mattermost-server/issues/11288
// rep--lace github.com/mattermost/mattermost-server v5.9.0+incompatible => /home/rmaneschi/projects/go_workspace/src/github.com/mattermost/mattermost-server

replace willnorris.com/go/imageproxy v0.8.1-0.20190326225038-d4246a08fdec => willnorris.com/go/imageproxy v0.8.1-0.20190422234945-d4246a08fdec
