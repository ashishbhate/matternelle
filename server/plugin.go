package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"reflect"
	"sync"

	pluginapi "github.com/mattermost/mattermost-plugin-api"
	"github.com/mattermost/mattermost-server/v6/model"
	"github.com/mattermost/mattermost-server/v6/plugin"
	"github.com/pkg/errors"
)

type Plugin struct {
	plugin.MattermostPlugin
	client *pluginapi.Client

	BotUserID    string
	Users        []*AppUser
	Applications []*App

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path; path {
	case "/ws":
		echo(w, r)
	default:
		http.NotFound(w, r)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	log.Print(reflect.TypeOf(w))
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (p *Plugin) OnActivate() error {
	p.client = pluginapi.NewClient(p.API, p.Driver)
	botId, err := p.createOrRetrieveBot()
	if err != nil {
		return errors.Wrap(err, "failed to create or retrieve Matternelle bot")
	}

	p.BotUserID = botId

	applications, err := p.GetApplicationsFromKVStore()
	if err != nil {
		return errors.Wrap(err, "failed to retrieve applications from KVStore")
	}

	p.Applications = applications

	return p.registerCommand()
}

func (p *Plugin) createOrRetrieveBot() (string, error) {
	botId, err := p.client.Bot.EnsureBot(&model.Bot{
		Username:    "matternelle",
		DisplayName: "Matternelle",
		Description: "A bot account created by the plugin Matternelle",
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to ensure bot")
	}
	bundlePath, err := p.API.GetBundlePath()
	if err != nil {
		return "", errors.Wrap(err, "failed to retrieve bundle path")
	}
	profileImage, err := ioutil.ReadFile(filepath.Join(bundlePath, "assets", "profile.jpg"))
	if err != nil {
		return "", errors.Wrap(err, "failed to read profile image")
	}
	if err := p.API.SetProfileImage(botId, profileImage); err != nil {
		return "", errors.Wrap(err, "failed to set profile image")
	}
	return botId, nil
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
