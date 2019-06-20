package main

func (p *Plugin) NewAppUser(u *AppUser) error {
	return p.PostPluginMessage("New app user connected")
}

func (p *Plugin) AppUserLeave(u *AppUser) error {
	return p.PostPluginMessage("App user disconnected")
}

func (p *Plugin) NewMessageFromAppUser(msg string) error {
	return nil
}
