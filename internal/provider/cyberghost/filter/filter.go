package filter

import (
	"errors"

	"github.com/qdm12/gluetun/internal/configuration"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/utils"
)

var ErrGroupMismatchesProtocol = errors.New("server group does not match protocol")

func Servers(allServers []models.CyberghostServer,
	selection configuration.ServerSelection) (
	servers []models.CyberghostServer, err error) {
	if len(selection.Groups) == 0 {
		if selection.OpenVPN.TCP {
			selection.Groups = tcpGroupChoices(allServers)
		} else {
			selection.Groups = udpGroupChoices(allServers)
		}
	}
	// Check each group match the protocol
	groupsCheckFn := groupsAreAllUDP
	if selection.OpenVPN.TCP {
		groupsCheckFn = groupsAreAllTCP
	}
	if err := groupsCheckFn(selection.Groups); err != nil {
		return nil, err
	}

	for _, server := range allServers {
		switch {
		case
			utils.FilterByPossibilities(server.Group, selection.Groups),
			utils.FilterByPossibilities(server.Region, selection.Regions),
			utils.FilterByPossibilities(server.Hostname, selection.Hostnames):
		default:
			servers = append(servers, server)
		}
	}

	if len(servers) == 0 {
		return nil, utils.NoServerFoundError(selection)
	}

	return servers, nil
}
