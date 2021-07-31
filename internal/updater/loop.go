package updater

import (
	"net/http"
	"time"

	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/loopstate"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/storage"
	"github.com/qdm12/gluetun/internal/updater/state"
	"github.com/qdm12/golibs/logging"
)

var _ Looper = (*Loop)(nil)

type Looper interface {
	RestartTickerRunner
	Runner
	loopstate.Getter
	loopstate.Applier
	SettingsGetSetter
}

type Loop struct {
	statusManager loopstate.Manager
	state         state.Manager
	// Objects
	updater       ServerUpdater
	storage       storage.Storage
	setAllServers func(allServers models.AllServers)
	logger        logging.Logger
	// Internals
	start        chan struct{}
	running      chan models.LoopStatus
	stop         chan struct{}
	stopped      chan struct{}
	updateTicker chan struct{}
	backoffTime  time.Duration
	userTrigger  bool
	// Mock functions
	timeNow   func() time.Time
	timeSince func(time.Time) time.Duration
}

const defaultBackoffTime = 5 * time.Second

func NewLoop(settings configuration.Updater, currentServers models.AllServers,
	storage storage.Storage, setAllServers func(allServers models.AllServers),
	client *http.Client, logger logging.Logger) *Loop {
	start := make(chan struct{})
	running := make(chan models.LoopStatus)
	stop := make(chan struct{})
	stopped := make(chan struct{})
	updateTicker := make(chan struct{})

	statusManager := loopstate.New(constants.Stopped, start, running, stop, stopped)
	state := state.New(statusManager, settings, updateTicker)

	return &Loop{
		statusManager: statusManager,
		state:         state,
		updater:       New(settings, client, currentServers, logger),
		storage:       storage,
		setAllServers: setAllServers,
		logger:        logger,
		start:         start,
		running:       running,
		stop:          stop,
		stopped:       stopped,
		updateTicker:  updateTicker,
		userTrigger:   true,
		timeNow:       time.Now,
		timeSince:     time.Since,
		backoffTime:   defaultBackoffTime,
	}
}
