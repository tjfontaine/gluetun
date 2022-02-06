package storage

import (
	"fmt"
	"reflect"

	"github.com/qdm12/gluetun/internal/models"
)

// GetServers returns all the servers merged from the
// persisted file and from the binary built-in servers data.
func (s *Storage) GetServers(filepath string) (servers models.AllServers, err error) {
	// error returned covered by unit test
	hardcoded, _ := parseHardcodedServers()

	persisted, err := s.readFromFile(filepath, hardcoded)
	if err != nil {
		return servers, fmt.Errorf("cannot read servers from file: %w", err)
	}

	hardcodedCount := countServers(hardcoded)
	countOnFile := countServers(persisted)

	var merged models.AllServers
	if countOnFile == 0 {
		s.logger.Info(fmt.Sprintf(
			"creating %s using %d hardcoded servers",
			filepath, hardcodedCount))
		merged = hardcoded
	} else {
		s.logger.Info(fmt.Sprintf(
			"merging by most recent %d hardcoded servers and %d servers read from file",
			hardcodedCount, countOnFile))
		merged = s.mergeServers(hardcoded, persisted)
	}

	// Write file if there is a difference between the
	// merged servers and the persisted servers.
	if !reflect.DeepEqual(persisted, merged) {
		err = flushToFile(filepath, merged)
		if err != nil {
			return servers, fmt.Errorf("cannot flush to file: %w", err)
		}
	}

	return merged, nil
}
