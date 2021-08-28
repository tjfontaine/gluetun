package constants

import (
	"sort"

	"github.com/qdm12/gluetun/internal/models"
)

func CyberghostRegionChoices(servers []models.CyberghostServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Region
	}
	return makeUnique(choices)
}

func CyberghostGroupChoices(servers []models.CyberghostServer) (choices []string) {
	uniqueChoices := map[string]struct{}{}
	for _, server := range servers {
		uniqueChoices[server.Group] = struct{}{}
	}

	choices = make([]string, 0, len(uniqueChoices))
	for choice := range uniqueChoices {
		choices = append(choices, choice)
	}

	sortable := sort.StringSlice(choices)
	sortable.Sort()

	return sortable
}

func CyberghostHostnameChoices(servers []models.CyberghostServer) (choices []string) {
	choices = make([]string, len(servers))
	for i := range servers {
		choices[i] = servers[i].Hostname
	}
	return makeUnique(choices)
}
