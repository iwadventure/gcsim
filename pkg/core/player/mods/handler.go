package mods

import (
	"github.com/genshinsim/gcsim/pkg/core/glog"
)

const MaxTeamSize = 4

type Handler struct {
	f     *int
	log   glog.Logger
	debug bool

	shieldBonusMods [MaxTeamSize][]shieldBonusMod

	// team []char
}

// type char struct {
// 	stats [attributes.EndStatType]float64
// 	atk   float64 //base atk including weapon
// 	hp    float64
// 	def   float64
// }
