package configuration

import (
	"fmt"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/golibs/params"
)

func (settings *Provider) readTorguard(r reader) (err error) {
	settings.Name = constants.Torguard
	servers := r.servers.GetTorguard()

	settings.ServerSelection.TargetIP, err = readTargetIP(r.env)
	if err != nil {
		return err
	}

	settings.ServerSelection.Countries, err = r.env.CSVInside("COUNTRY", constants.TorguardCountryChoices(servers))
	if err != nil {
		return fmt.Errorf("environment variable COUNTRY: %w", err)
	}

	settings.ServerSelection.Cities, err = r.env.CSVInside("CITY", constants.TorguardCityChoices(servers))
	if err != nil {
		return fmt.Errorf("environment variable CITY: %w", err)
	}

	settings.ServerSelection.Hostnames, err = r.env.CSVInside("SERVER_HOSTNAME",
		constants.TorguardHostnameChoices(servers))
	if err != nil {
		return fmt.Errorf("environment variable SERVER_HOSTNAME: %w", err)
	}

	if settings.ServerSelection.VPN == constants.OpenVPN {
		err = settings.ServerSelection.OpenVPN.readProtocolAndPort(r.env)
	} else {
		err = settings.ServerSelection.Wireguard.readTorguard(r.env)
	}

	return err
}

func (settings *WireguardSelection) readTorguard(env params.Interface) (err error) {
	settings.PublicKey, err = env.Get("WIREGUARD_PUBLIC_KEY",
		params.CaseSensitiveValue(), params.Compulsory())
	if err != nil {
		return fmt.Errorf("environment variable WIREGUARD_PUBLIC_KEY: %w", err)
	}

	settings.EndpointIP, err = readWireguardEndpointIP(env)
	if err != nil {
		return err
	}

	const portCompulsory = true
	settings.EndpointPort, err = readWireguardEndpointPort(env, nil, portCompulsory)
	if err != nil {
		return err
	}

	return nil
}
