package main

import (
	"sync"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type Plugin struct {
	plugin.MattermostPlugin
	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex
	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// UserHasBeenCreated is invoked after a user was created.
func (p *Plugin) UserHasBeenCreated(c *plugin.Context, user *model.User) {
	theme := p.getConfiguration().CustomTheme
	if theme == "" {
		return
	}

	pref := model.Preference{
		UserId:   user.Id,
		Category: "theme",
		Name:     "custom", // Указываем имя кастомной темы
		Value:    theme,
	}
	prefs := []model.Preference{pref}

	if appErr := p.API.UpdatePreferencesForUser(user.Id, prefs); appErr != nil {
		p.API.LogError("Failed to set custom theme for user", "user_id", user.Id, "error", appErr.Error())
	}
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
