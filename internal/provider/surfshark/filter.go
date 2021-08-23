package surfshark

import (
	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (s *Surfshark) filterServers(selection configuration.ServerSelection) (
	servers []models.SurfsharkServer, err error) {
	wantOpenVPN := selection.VPN == constants.OpenVPN
	wantWireguard := selection.VPN == constants.Wireguard
	for _, server := range s.servers {
		switch {
		case
			utils.FilterByPossibilities(server.Region, selection.Regions),
			utils.FilterByPossibilities(server.Country, selection.Countries),
			utils.FilterByPossibilities(server.City, selection.Cities),
			utils.FilterByPossibilities(server.Hostname, selection.Hostnames),
			wantOpenVPN && !server.OpenVPN,
			wantWireguard && !server.Wireguard,
			utils.FilterByProtocol(selection, server.TCP, server.UDP),
			selection.MultiHopOnly && !server.MultiHop:
		default:
			servers = append(servers, server)
		}
	}

	if len(servers) == 0 {
		return nil, utils.NoServerFoundError(selection)
	}

	return servers, nil
}
