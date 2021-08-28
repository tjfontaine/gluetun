package storage

import (
	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/cyberghost/constants"
	utils "github.com/qdm12/gluetun/internal/provider/utils/storage"
)

func Merge(hardcoded, persisted models.CyberghostServers,
	logger utils.Logger) (merged models.CyberghostServers) {
	if persisted.Timestamp <= hardcoded.Timestamp {
		return hardcoded
	}

	versionDiff := int(hardcoded.Version) - int(persisted.Version)
	if versionDiff > 0 {
		logger.Info(utils.DiffVersionMsg(constants.Name, versionDiff))
		return hardcoded
	}

	logger.Info(utils.DiffTimeMsg(constants.Name, persisted.Timestamp, hardcoded.Timestamp))
	return persisted
}
