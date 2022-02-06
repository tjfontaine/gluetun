package storage

import (
	"github.com/qdm12/gluetun/internal/models"
)

func countServers(allServers models.AllServers) int {
	return len(allServers.Cyberghost.Servers) +
		len(allServers.Expressvpn.Servers) +
		len(allServers.Fastestvpn.Servers) +
		len(allServers.HideMyAss.Servers) +
		len(allServers.Ipvanish.Servers) +
		len(allServers.Ivpn.Servers) +
		len(allServers.Mullvad.Servers) +
		len(allServers.Nordvpn.Servers) +
		len(allServers.Perfectprivacy.Servers) +
		len(allServers.Privado.Servers) +
		len(allServers.Pia.Servers) +
		len(allServers.Privatevpn.Servers) +
		len(allServers.Protonvpn.Servers) +
		len(allServers.Purevpn.Servers) +
		len(allServers.Surfshark.Servers) +
		len(allServers.Torguard.Servers) +
		len(allServers.VPNUnlimited.Servers) +
		len(allServers.Vyprvpn.Servers) +
		len(allServers.Wevpn.Servers) +
		len(allServers.Windscribe.Servers)
}
