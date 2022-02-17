package vpnsecure

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/qdm12/gluetun/internal/constants/vpn"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/updater/resolver"
)

var ErrNotEnoughServers = errors.New("not enough servers found")

func GetServers(ctx context.Context, client *http.Client,
	presolver resolver.Parallel, minServers int) (
	servers []models.Server, warnings []string, err error) {
	servers, err = fetchServers(ctx, client)
	if err != nil {
		return nil, nil, fmt.Errorf("cannot fetch servers: %w", err)
	}

	if len(servers) < minServers {
		return nil, nil, fmt.Errorf("%w: %d and expected at least %d",
			ErrNotEnoughServers, len(servers), minServers)
	}

	hts := make(hostToServer, len(servers))
	for _, server := range servers {
		hts[server.Hostname] = server
	}

	hosts := hts.toHostsSlice()

	hostToIPs, newWarnings, err := resolveHosts(ctx, presolver, hosts, minServers)
	warnings = append(warnings, newWarnings...)
	if err != nil {
		return nil, warnings, err
	}

	hts.adaptWithIPs(hostToIPs)

	servers = hts.toServersSlice()

	if len(servers) < minServers {
		return nil, warnings, fmt.Errorf("%w: %d and expected at least %d",
			ErrNotEnoughServers, len(servers), minServers)
	}

	for i := range servers {
		servers[i].VPN = vpn.OpenVPN
		servers[i].UDP = true
		servers[i].TCP = true
	}

	sortServers(servers)

	return servers, warnings, nil
}
