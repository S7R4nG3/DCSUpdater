package flags

import (
	"github.com/S7R4nG3/DCSUpdater/types"
)

func Get(s types.Version) []string {
	var (
		DcsStable    = []string{"update", "@release"}
		DcsOpenBeta  = []string{"update", "@openbeta"}
		DcsOpenAlpha = []string{"update", "@openalpha"}
		Repair       = []string{"repair"}
		Cleanup      = []string{"cleanup"}
	)

	switch s {
	case types.Stable:
		return DcsStable
	case types.OpenBeta:
		return DcsOpenBeta
	case types.OpenAlpha:
		return DcsOpenAlpha
	case types.Repair:
		return Repair
	case types.Cleanup:
		return Cleanup
	default:
		return []string{}
	}
}
