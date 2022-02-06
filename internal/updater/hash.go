package updater

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
)

type makeServersDataFunc func(servers models.AllServers) (data serversData, err error)

func makeServersData(servers models.AllServers) (vpnProviderData map[string]serversData, err error) {
	vpnProviderToFunc := map[string]makeServersDataFunc{
		constants.Cyberghost: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Cyberghost.Servers)
			data.serversHexDigest, err = hashServers(servers.Cyberghost.Servers)
			return data, err
		},
		constants.Expressvpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Expressvpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Expressvpn.Servers)
			return data, err
		},
		constants.Fastestvpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Fastestvpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Fastestvpn.Servers)
			return data, err
		},
		constants.HideMyAss: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.HideMyAss.Servers)
			data.serversHexDigest, err = hashServers(servers.HideMyAss.Servers)
			return data, err
		},
		constants.Ipvanish: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Ipvanish.Servers)
			data.serversHexDigest, err = hashServers(servers.Ipvanish.Servers)
			return data, err
		},
		constants.Ivpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Ivpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Ivpn.Servers)
			return data, err
		},
		constants.Mullvad: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Mullvad.Servers)
			data.serversHexDigest, err = hashServers(servers.Mullvad.Servers)
			return data, err
		},
		constants.Nordvpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Nordvpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Nordvpn.Servers)
			return data, err
		},
		constants.Perfectprivacy: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Perfectprivacy.Servers)
			data.serversHexDigest, err = hashServers(servers.Perfectprivacy.Servers)
			return data, err
		},
		constants.Privado: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Privado.Servers)
			data.serversHexDigest, err = hashServers(servers.Privado.Servers)
			return data, err
		},
		constants.PrivateInternetAccess: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Pia.Servers)
			data.serversHexDigest, err = hashServers(servers.Pia.Servers)
			return data, err
		},
		constants.Privatevpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Privatevpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Privatevpn.Servers)
			return data, err
		},
		constants.Protonvpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Protonvpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Protonvpn.Servers)
			return data, err
		},
		constants.Purevpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Purevpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Purevpn.Servers)
			return data, err
		},
		constants.Surfshark: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Surfshark.Servers)
			data.serversHexDigest, err = hashServers(servers.Surfshark.Servers)
			return data, err
		},
		constants.Torguard: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Torguard.Servers)
			data.serversHexDigest, err = hashServers(servers.Torguard.Servers)
			return data, err
		},
		constants.VPNUnlimited: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.VPNUnlimited.Servers)
			data.serversHexDigest, err = hashServers(servers.VPNUnlimited.Servers)
			return data, err
		},
		constants.Vyprvpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Vyprvpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Vyprvpn.Servers)
			return data, err
		},
		constants.Wevpn: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Wevpn.Servers)
			data.serversHexDigest, err = hashServers(servers.Wevpn.Servers)
			return data, err
		},
		constants.Windscribe: func(servers models.AllServers) (data serversData, err error) {
			data.serversCount = len(servers.Windscribe.Servers)
			data.serversHexDigest, err = hashServers(servers.Windscribe.Servers)
			return data, err
		},
	}

	vpnProviderData = make(map[string]serversData, len(vpnProviderToFunc))
	for provider, f := range vpnProviderToFunc {
		vpnProviderData[provider], err = f(servers)
		if err != nil {
			return nil, fmt.Errorf("for provider %s: %w", provider, err)
		}
	}

	return vpnProviderData, nil
}

func hashServers(servers interface{}) (hexDigest string, err error) {
	b, err := json.Marshal(servers)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(b)

	hexDigest = hex.EncodeToString(sum[:])
	return hexDigest, nil
}
