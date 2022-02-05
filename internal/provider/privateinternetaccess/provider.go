package privateinternetaccess

import (
	"math/rand"
	"path/filepath"
	"time"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
)

type PIA struct {
	servers    []models.PIAServer
	randSource rand.Source
	timeNow    func() time.Time
	// Port forwarding
	portForwardPath string
	authFilePath    string
}

func New(servers []models.PIAServer, randSource rand.Source,
	gluetunDir string, timeNow func() time.Time) *PIA {
	const jsonPortForwardFilename = "piaportforward.json"
	jsonPortForwardPath := filepath.Join(gluetunDir, jsonPortForwardFilename)
	return &PIA{
		servers:         servers,
		timeNow:         timeNow,
		randSource:      randSource,
		portForwardPath: jsonPortForwardPath,
		authFilePath:    constants.OpenVPNAuthConf,
	}
}
