package main

func (p *Plugin) NewAppUser(u *AppUser) error {
	return nil
}

func (p *Plugin) AppUserLeave(u *AppUser) error {
	return nil
}

func (p *Plugin) NewMessageFromAppUser(msg string) error {
	return nil
}
