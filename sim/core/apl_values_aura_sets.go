package core

import (
	"fmt"

	"github.com/wowsims/cata/sim/core/proto"
	"github.com/wowsims/cata/sim/core/stats"
)

type APLValueAllTrinketStatProcsActive struct {
	DefaultAPLValueImpl

	statTypesToMatch []stats.Stat
	matchingAuras    []*StatBuffAura
}

func (rot *APLRotation) newValueAllTrinketStatProcsActive(config *proto.APLValueAllTrinketStatProcsActive) APLValue {
	statTypesToMatch := stats.IntTupleToStatsList(config.StatType1, config.StatType2)
	matchingAuras := rot.GetAPLTrinketProcAuras(statTypesToMatch)

	if len(matchingAuras) == 0 {
		return nil
	}

	return &APLValueAllTrinketStatProcsActive{
		statTypesToMatch: statTypesToMatch,
		matchingAuras:    matchingAuras,
	}
}
func (value *APLValueAllTrinketStatProcsActive) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueAllTrinketStatProcsActive) GetBool(sim *Simulation) bool {
	for _, aura := range value.matchingAuras {
		if !aura.IsActive() || (aura.GetStacks() < aura.MaxStacks) {
			return false
		}
	}

	return true
}
func (value *APLValueAllTrinketStatProcsActive) String() string {
	return fmt.Sprintf("All Trinket Stat Procs Active(%s)", StringFromStatTypes(value.statTypesToMatch))
}

type APLValueAnyTrinketStatProcsActive struct {
	DefaultAPLValueImpl

	statTypesToMatch []stats.Stat
	matchingAuras    []*StatBuffAura
}

func (rot *APLRotation) newValueAnyTrinketStatProcsActive(config *proto.APLValueAnyTrinketStatProcsActive) APLValue {
	statTypesToMatch := stats.IntTupleToStatsList(config.StatType1, config.StatType2)
	matchingAuras := rot.GetAPLTrinketProcAuras(statTypesToMatch)

	if len(matchingAuras) == 0 {
		return nil
	}

	return &APLValueAnyTrinketStatProcsActive{
		statTypesToMatch: statTypesToMatch,
		matchingAuras:    matchingAuras,
	}
}
func (value *APLValueAnyTrinketStatProcsActive) Type() proto.APLValueType {
	return proto.APLValueType_ValueTypeBool
}
func (value *APLValueAnyTrinketStatProcsActive) GetBool(sim *Simulation) bool {
	for _, aura := range value.matchingAuras {
		if aura.IsActive() && (aura.GetStacks() == aura.MaxStacks) {
			return true
		}
	}

	return false
}
func (value *APLValueAnyTrinketStatProcsActive) String() string {
	return fmt.Sprintf("Any Trinket Stat Procs Active(%s)", StringFromStatTypes(value.statTypesToMatch))
}
