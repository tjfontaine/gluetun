package updater

import (
	"context"
	"fmt"

	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/updater/providers/cyberghost"
	"github.com/qdm12/gluetun/internal/updater/providers/expressvpn"
	"github.com/qdm12/gluetun/internal/updater/providers/fastestvpn"
	"github.com/qdm12/gluetun/internal/updater/providers/hidemyass"
	"github.com/qdm12/gluetun/internal/updater/providers/ipvanish"
	"github.com/qdm12/gluetun/internal/updater/providers/ivpn"
	"github.com/qdm12/gluetun/internal/updater/providers/mullvad"
	"github.com/qdm12/gluetun/internal/updater/providers/nordvpn"
	"github.com/qdm12/gluetun/internal/updater/providers/perfectprivacy"
	"github.com/qdm12/gluetun/internal/updater/providers/pia"
	"github.com/qdm12/gluetun/internal/updater/providers/privado"
	"github.com/qdm12/gluetun/internal/updater/providers/privatevpn"
	"github.com/qdm12/gluetun/internal/updater/providers/protonvpn"
	"github.com/qdm12/gluetun/internal/updater/providers/purevpn"
	"github.com/qdm12/gluetun/internal/updater/providers/surfshark"
	"github.com/qdm12/gluetun/internal/updater/providers/torguard"
	"github.com/qdm12/gluetun/internal/updater/providers/vpnunlimited"
	"github.com/qdm12/gluetun/internal/updater/providers/vyprvpn"
	"github.com/qdm12/gluetun/internal/updater/providers/wevpn"
	"github.com/qdm12/gluetun/internal/updater/providers/windscribe"
)

func (u *Updater) updateCyberghost(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Cyberghost])
	servers, err := cyberghost.GetServers(ctx, u.presolver, minServers)
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Cyberghost, servers) {
		return nil
	}

	allServers.Cyberghost.Timestamp = u.timeNow().Unix()
	allServers.Cyberghost.Servers = servers
	return nil
}

func (u *Updater) updateExpressvpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Expressvpn])
	servers, warnings, err := expressvpn.GetServers(
		ctx, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("ExpressVPN: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Expressvpn, servers) {
		return nil
	}

	allServers.Expressvpn.Timestamp = u.timeNow().Unix()
	allServers.Expressvpn.Servers = servers
	return nil
}

func (u *Updater) updateFastestvpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Fastestvpn])
	servers, warnings, err := fastestvpn.GetServers(
		ctx, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("FastestVPN: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Fastestvpn, servers) {
		return nil
	}

	allServers.Fastestvpn.Timestamp = u.timeNow().Unix()
	allServers.Fastestvpn.Servers = servers
	return nil
}

func (u *Updater) updateHideMyAss(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.HideMyAss])
	servers, warnings, err := hidemyass.GetServers(
		ctx, u.client, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("HideMyAss: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.HideMyAss, servers) {
		return nil
	}

	allServers.HideMyAss.Timestamp = u.timeNow().Unix()
	allServers.HideMyAss.Servers = servers
	return nil
}

func (u *Updater) updateIpvanish(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Ipvanish])
	servers, warnings, err := ipvanish.GetServers(
		ctx, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("Ipvanish: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Ipvanish, servers) {
		return nil
	}

	allServers.Ipvanish.Timestamp = u.timeNow().Unix()
	allServers.Ipvanish.Servers = servers
	return nil
}

func (u *Updater) updateIvpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Ivpn])
	servers, warnings, err := ivpn.GetServers(
		ctx, u.client, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("Ivpn: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Ivpn, servers) {
		return nil
	}

	allServers.Ivpn.Timestamp = u.timeNow().Unix()
	allServers.Ivpn.Servers = servers
	return nil
}

func (u *Updater) updateMullvad(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Mullvad])
	servers, err := mullvad.GetServers(ctx, u.client, minServers)
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Mullvad, servers) {
		return nil
	}

	allServers.Mullvad.Timestamp = u.timeNow().Unix()
	allServers.Mullvad.Servers = servers
	return nil
}

func (u *Updater) updateNordvpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Nordvpn])
	servers, warnings, err := nordvpn.GetServers(ctx, u.client, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("NordVPN: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Nordvpn, servers) {
		return nil
	}

	allServers.Nordvpn.Timestamp = u.timeNow().Unix()
	allServers.Nordvpn.Servers = servers
	return nil
}

func (u *Updater) updatePerfectprivacy(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Perfectprivacy])
	servers, warnings, err := perfectprivacy.GetServers(ctx, u.unzipper, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn(constants.Perfectprivacy + ": " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Perfectprivacy, servers) {
		return nil
	}

	allServers.Perfectprivacy.Timestamp = u.timeNow().Unix()
	allServers.Perfectprivacy.Servers = servers
	return nil
}

func (u *Updater) updatePIA(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.PrivateInternetAccess])
	servers, err := pia.GetServers(ctx, u.client, minServers)
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.PrivateInternetAccess, servers) {
		return nil
	}

	allServers.Pia.Timestamp = u.timeNow().Unix()
	allServers.Pia.Servers = servers
	return nil
}

func (u *Updater) updatePrivado(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Privado])
	servers, warnings, err := privado.GetServers(
		ctx, u.unzipper, u.client, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("Privado: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Privado, servers) {
		return nil
	}

	allServers.Privado.Timestamp = u.timeNow().Unix()
	allServers.Privado.Servers = servers
	return nil
}

func (u *Updater) updatePrivatevpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Privatevpn])
	servers, warnings, err := privatevpn.GetServers(
		ctx, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("PrivateVPN: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Privatevpn, servers) {
		return nil
	}

	allServers.Privatevpn.Timestamp = u.timeNow().Unix()
	allServers.Privatevpn.Servers = servers
	return nil
}

func (u *Updater) updateProtonvpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Privatevpn])
	servers, warnings, err := protonvpn.GetServers(ctx, u.client, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("ProtonVPN: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Protonvpn, servers) {
		return nil
	}

	allServers.Protonvpn.Timestamp = u.timeNow().Unix()
	allServers.Protonvpn.Servers = servers
	return nil
}

func (u *Updater) updatePurevpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Purevpn])
	servers, warnings, err := purevpn.GetServers(
		ctx, u.client, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("PureVPN: " + warning)
		}
	}
	if err != nil {
		return fmt.Errorf("cannot update Purevpn servers: %w", err)
	}

	if !u.didServersChange(constants.Purevpn, servers) {
		return nil
	}

	allServers.Purevpn.Timestamp = u.timeNow().Unix()
	allServers.Purevpn.Servers = servers
	return nil
}

func (u *Updater) updateSurfshark(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Surfshark])
	servers, warnings, err := surfshark.GetServers(
		ctx, u.unzipper, u.client, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("Surfshark: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Surfshark, servers) {
		return nil
	}

	allServers.Surfshark.Timestamp = u.timeNow().Unix()
	allServers.Surfshark.Servers = servers
	return nil
}

func (u *Updater) updateTorguard(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Torguard])
	servers, warnings, err := torguard.GetServers(
		ctx, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("Torguard: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Torguard, servers) {
		return nil
	}

	allServers.Torguard.Timestamp = u.timeNow().Unix()
	allServers.Torguard.Servers = servers
	return nil
}

func (u *Updater) updateVPNUnlimited(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.VPNUnlimited])
	servers, warnings, err := vpnunlimited.GetServers(
		ctx, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn(constants.VPNUnlimited + ": " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.VPNUnlimited, servers) {
		return nil
	}

	allServers.VPNUnlimited.Timestamp = u.timeNow().Unix()
	allServers.VPNUnlimited.Servers = servers
	return nil
}

func (u *Updater) updateVyprvpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Vyprvpn])
	servers, warnings, err := vyprvpn.GetServers(
		ctx, u.unzipper, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("VyprVPN: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Vyprvpn, servers) {
		return nil
	}

	allServers.Vyprvpn.Timestamp = u.timeNow().Unix()
	allServers.Vyprvpn.Servers = servers
	return nil
}

func (u *Updater) updateWevpn(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Wevpn])
	servers, warnings, err := wevpn.GetServers(ctx, u.presolver, minServers)
	if *u.options.CLI {
		for _, warning := range warnings {
			u.logger.Warn("WeVPN: " + warning)
		}
	}
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Wevpn, servers) {
		return nil
	}

	allServers.Wevpn.Timestamp = u.timeNow().Unix()
	allServers.Wevpn.Servers = servers
	return nil
}

func (u *Updater) updateWindscribe(ctx context.Context, allServers *models.AllServers) (err error) {
	minServers := getMinServers(u.vpnProviderData[constants.Windscribe])
	servers, err := windscribe.GetServers(ctx, u.client, minServers)
	if err != nil {
		return err
	}

	if !u.didServersChange(constants.Windscribe, servers) {
		return nil
	}

	allServers.Windscribe.Timestamp = u.timeNow().Unix()
	allServers.Windscribe.Servers = servers
	return nil
}

func getMinServers(serversData serversData) (minServers int) {
	serversCount := serversData.serversCount
	const minRatio = 0.8
	return int(minRatio * float64(serversCount))
}

func (u *Updater) didServersChange(vpnProvider string, servers interface{}) (changed bool) {
	previousHexDigest := u.vpnProviderData[vpnProvider].serversHexDigest
	currentHexDigest, err := hashServers(servers)
	if err != nil {
		panic(err)
	}
	return previousHexDigest != currentHexDigest
}
