package state

import (
	"sync"

	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/loopstate"
)

var _ Manager = (*State)(nil)

type Manager interface {
	SettingsGetSetter
}

func New(statusApplier loopstate.Applier,
	settings configuration.Updater,
	updateTicker chan<- struct{}) *State {
	return &State{
		statusApplier: statusApplier,
		updateTicker:  updateTicker,
		settings:      settings,
	}
}

type State struct {
	statusApplier loopstate.Applier
	updateTicker  chan<- struct{}

	settings   configuration.Updater
	settingsMu sync.RWMutex
}
