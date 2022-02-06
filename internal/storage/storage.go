// Package storage defines interfaces to interact with the files persisted such as the list of servers.
package storage

import (
	"fmt"

	"github.com/qdm12/gluetun/internal/models"
)

//go:generate mockgen -destination=infoer_mock_test.go -package $GOPACKAGE . Infoer

type Storage struct {
	logger Infoer
}

func New(logger Infoer) (storage *Storage) {
	return &Storage{
		logger: logger,
	}
}

func (s *Storage) Init(serversFilepath string) (servers models.AllServers, err error) {
	servers, err = s.GetServers(serversFilepath)
	if err != nil {
		return servers, fmt.Errorf("cannot get servers: %w", err)
	}

	s.logger = newNoopInfoer()

	return servers, nil
}
