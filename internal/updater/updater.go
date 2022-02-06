// Package updater implements update mechanisms for each VPN provider servers.
package updater

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/qdm12/gluetun/internal/configuration/settings"
	"github.com/qdm12/gluetun/internal/constants"
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/storage"
	"github.com/qdm12/gluetun/internal/updater/resolver"
	"github.com/qdm12/gluetun/internal/updater/unzip"
)

type ServerUpdater interface {
	UpdateServers(ctx context.Context) (allServers models.AllServers, err error)
}

type Updater struct {
	// configuration
	options settings.Updater

	// state
	vpnProviderData map[string]serversData
	storage         storage.Storage // TODO interface

	// Functions for tests
	logger    Logger
	timeNow   func() time.Time
	presolver resolver.Parallel
	client    *http.Client
	unzipper  unzip.Unzipper
}

type serversData struct {
	serversHexDigest string
	serversCount     int
}

func New(settings settings.Updater, httpClient *http.Client,
	currentServers models.AllServers, logger Logger) (u *Updater, err error) {
	unzipper := unzip.New(httpClient)
	presolver := resolver.NewParallelResolver(settings.DNSAddress.String())

	vpnProviderData, err := makeServersData(currentServers)
	if err != nil {
		return nil, err
	}

	return &Updater{
		logger:          logger,
		timeNow:         time.Now,
		presolver:       presolver,
		client:          httpClient,
		unzipper:        unzipper,
		options:         settings,
		vpnProviderData: vpnProviderData,
	}, nil
}

type updateFunc func(ctx context.Context, allServers *models.AllServers) (err error)

func (u *Updater) UpdateServers(ctx context.Context) (
	allServers models.AllServers, err error) {
	allServers, err = u.storage.GetServers(constants.ServersData)
	if err != nil {
		return allServers, err
	}

	for _, provider := range u.options.Providers {
		u.logger.Info("updating " + strings.Title(provider) + " servers...")
		updateProvider := u.getUpdateFunction(provider)

		// TODO support servers offering only TCP or only UDP
		// for NordVPN and PureVPN
		err = updateProvider(ctx, &allServers)
		if err != nil {
			if ctxErr := ctx.Err(); ctxErr != nil {
				return allServers, ctxErr
			}
			u.logger.Error(err.Error())
		}
	}

	return allServers, nil
}

func (u *Updater) getUpdateFunction(provider string) (updateFunction updateFunc) {
	switch provider {
	case constants.Custom:
		panic("custom provider is not meant to be updated")
	case constants.Cyberghost:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateCyberghost(ctx, allServers)
		}
	case constants.Expressvpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateExpressvpn(ctx, allServers)
		}
	case constants.Fastestvpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateFastestvpn(ctx, allServers)
		}
	case constants.HideMyAss:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateHideMyAss(ctx, allServers)
		}
	case constants.Ipvanish:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateIpvanish(ctx, allServers)
		}
	case constants.Ivpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateIvpn(ctx, allServers)
		}
	case constants.Mullvad:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateMullvad(ctx, allServers)
		}
	case constants.Nordvpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateNordvpn(ctx, allServers)
		}
	case constants.Perfectprivacy:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updatePerfectprivacy(ctx, allServers)
		}
	case constants.Privado:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updatePrivado(ctx, allServers)
		}
	case constants.PrivateInternetAccess:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updatePIA(ctx, allServers)
		}
	case constants.Privatevpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updatePrivatevpn(ctx, allServers)
		}
	case constants.Protonvpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateProtonvpn(ctx, allServers)
		}
	case constants.Purevpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updatePurevpn(ctx, allServers)
		}
	case constants.Surfshark:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateSurfshark(ctx, allServers)
		}
	case constants.Torguard:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateTorguard(ctx, allServers)
		}
	case constants.VPNUnlimited:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateVPNUnlimited(ctx, allServers)
		}
	case constants.Vyprvpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateVyprvpn(ctx, allServers)
		}
	case constants.Wevpn:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateWevpn(ctx, allServers)
		}
	case constants.Windscribe:
		return func(ctx context.Context, allServers *models.AllServers) (err error) {
			return u.updateWindscribe(ctx, allServers)
		}
	default:
		panic("provider " + provider + " is unknown")
	}
}
