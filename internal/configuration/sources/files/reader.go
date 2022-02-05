package files

import (
	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/configuration/sources"
)

var _ sources.Source = (*Reader)(nil)

type Reader struct {
	gluetunDir string
}

func New(gluetunDir string) *Reader {
	return &Reader{
		gluetunDir: gluetunDir,
	}
}

func (r *Reader) Read() (settings settings.Settings, err error) {
	settings.VPN, err = r.readVPN(r.gluetunDir)
	if err != nil {
		return settings, err
	}

	settings.System, err = r.readSystem()
	if err != nil {
		return settings, err
	}

	return settings, nil
}
