package filter

import (
	"fmt"
	"strings"

	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/provider/cyberghost/constants"
)

func tcpGroupChoices(servers []models.CyberghostServer) (choices []string) {
	const tcp = true
	return groupsForTCP(servers, tcp)
}

func udpGroupChoices(servers []models.CyberghostServer) (choices []string) {
	const tcp = false
	return groupsForTCP(servers, tcp)
}

func groupsForTCP(servers []models.CyberghostServer, tcp bool) (choices []string) {
	allGroups := constants.CyberghostGroupChoices(servers)
	choices = make([]string, 0, len(allGroups))
	for _, group := range allGroups {
		switch {
		case tcp && groupIsTCP(group):
			choices = append(choices, group)
		case !tcp && !groupIsTCP(group):
			choices = append(choices, group)
		}
	}
	return choices
}

func groupIsTCP(group string) bool {
	return strings.Contains(strings.ToLower(group), "tcp")
}

func groupsAreAllTCP(groups []string) error {
	for _, group := range groups {
		if !groupIsTCP(group) {
			return fmt.Errorf("%w: group %s for protocol TCP",
				ErrGroupMismatchesProtocol, group)
		}
	}
	return nil
}

func groupsAreAllUDP(groups []string) error {
	for _, group := range groups {
		if groupIsTCP(group) {
			return fmt.Errorf("%w: group %s for protocol UDP",
				ErrGroupMismatchesProtocol, group)
		}
	}
	return nil
}
