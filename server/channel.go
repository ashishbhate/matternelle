package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const applicationsKey = "applications"

type App struct {
	Name      string
	Token     string
	ChannelId string
}

func NewApp(appName string, token string, channelId string) *App {
	return &App{
		Name:      appName,
		Token:     token,
		ChannelId: channelId,
	}
}

func (p *Plugin) initialize(appName string, channelId string) (*App, error) {
	for _, app := range p.Applications {
		if appName == app.Name {
			return app, nil
		}
	}
	app := NewApp(appName, uuid.New().String(), channelId)
	p.Applications = append(p.Applications, app)
	return app, p.StoreApplicationsToKVStore(p.Applications)
}

func (p *Plugin) remove(appName string) error {
	var apps []*App
	for _, app := range p.Applications {
		if appName != app.Name {
			apps = append(apps, app)
		}
	}
	p.Applications = apps
	return p.StoreApplicationsToKVStore(p.Applications)
}

func (p *Plugin) StoreApplicationsToKVStore(apps []*App) error {
	str, err := json.Marshal(apps)
	if err != nil {
		return errors.Wrap(err, "can't serialize applications")
	}
	return p.API.KVSet(applicationsKey, str)
}

func (p *Plugin) GetApplicationsFromKVStore() ([]*App, error) {
	applications, err := p.API.KVGet(applicationsKey)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("can't get key %s from KV Store", applicationsKey))
	}
	if applications != nil {
		apps := []*App{}
		err := json.Unmarshal(applications, &apps)
		return apps, err
	}
	return nil, nil
}
