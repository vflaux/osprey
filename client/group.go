package client

import (
	"sort"
)

// Group organizes the targetEntry targets
type Group struct {
	name      string
	isDefault bool
	targets   []Target
}

// IsDefault returns true if this is the default group in the configuration
func (g *Group) IsDefault() bool {
	return g.isDefault
}

// Targets returns the list of targets belonging to this group
func (g *Group) Targets() []Target {
	return sortTargets(g.targets)
}

// Name returns the name of the group
func (g *Group) Name() string {
	return g.name
}

//Contains returns true if it contains the target
func (g *Group) Contains(target Target) bool {
	for _, current := range g.targets {
		if target.name == current.name {
			return true
		}
	}
	return false
}

func sortGroups(groups []Group) []Group {
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].name < groups[j].name
	})
	return groups
}