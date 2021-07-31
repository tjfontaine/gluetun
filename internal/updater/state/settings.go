package state

import (
	"context"
	"reflect"

	"github.com/qdm12/gluetun/internal/configuration"
)

type SettingsGetSetter interface {
	GetSettings() (settings configuration.Updater)
	SetSettings(ctx context.Context, settings configuration.Updater) (
		outcome string)
}

func (s *State) GetSettings() (settings configuration.Updater) {
	s.settingsMu.RLock()
	defer s.settingsMu.RUnlock()
	return s.settings
}

func (s *State) SetSettings(ctx context.Context, settings configuration.Updater) (
	outcome string) {
	s.settingsMu.Lock()
	settingsUnchanged := reflect.DeepEqual(s.settings, settings)
	if settingsUnchanged {
		s.settingsMu.Unlock()
		return "settings left unchanged"
	}
	s.settings = settings
	s.settingsMu.Unlock()
	s.updateTicker <- struct{}{}
	return "settings updated"
}
