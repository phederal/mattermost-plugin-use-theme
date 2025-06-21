package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/mattermost/mattermost-plugin-starter-template/server/command"
	"github.com/mattermost/mattermost-plugin-starter-template/server/store/kvstore"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/mattermost/mattermost/server/public/pluginapi"
	"github.com/mattermost/mattermost/server/public/pluginapi/cluster"
	"github.com/pkg/errors"
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
		Category: model.PREFERENCE_CATEGORY_THEME,
		Name:     "",
		Value:    theme,
	}
	prefs := []model.Preference{pref}
	p.API.UpdatePreferencesForUser(user.Id, prefs)
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
