package surfshark

import (
	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

func (s *Surfshark) GetConnection(selection configuration.ServerSelection) (
	connection models.Connection, err error) {
	protocol := utils.GetProtocol(selection)
	port := getPort(selection)

	servers, err := s.filterServers(selection)
	if err != nil {
		return connection, err
	}

	var connections []models.Connection
	for _, server := range servers {
		for _, IP := range server.IPs {
			connection := models.Connection{
				Type:     selection.VPN,
				IP:       IP,
				Port:     port,
				Protocol: protocol,
				PubKey:   server.WgPubKey,
			}
			connections = append(connections, connection)
		}
	}

	return utils.PickConnection(connections, selection, s.randSource)
}

func getPort(selection configuration.ServerSelection) (port uint16) {
	const (
		defaultOpenVPNTCP = 1443
		defaultOpenVPNUDP = 1194
		defaultWireguard  = 51820
	)
	return utils.GetPort(selection, defaultOpenVPNTCP,
		defaultOpenVPNUDP, defaultWireguard)
}
