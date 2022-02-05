package files

import (
	"fmt"
	"path/filepath"

	"github.com/qdm12/gluetun/internal/configuration/settings"
)

const (
	// OpenVPNClientKeyRelPath is the OpenVPN client key relative path inside the
	// gluetun directory.
	OpenVPNClientKeyRelPath = "client.key"
	// OpenVPNClientCertificateRelPath is the OpenVPN client certificate relative
	// path inside the gluetun directory.
	OpenVPNClientCertificateRelPath = "client.crt"
)

func (r *Reader) readOpenVPN(gluetunDir string) (settings settings.OpenVPN, err error) {
	settings.ClientKey, err = ReadFromFile(filepath.Join(gluetunDir, OpenVPNClientKeyRelPath))
	if err != nil {
		return settings, fmt.Errorf("cannot read client key: %w", err)
	}

	settings.ClientCrt, err = ReadFromFile(filepath.Join(gluetunDir, OpenVPNClientCertificateRelPath))
	if err != nil {
		return settings, fmt.Errorf("cannot read client certificate: %w", err)
	}

	return settings, nil
}
