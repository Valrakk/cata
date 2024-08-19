package core

import (
	"fmt"
	"time"

	"github.com/wowsims/cata/sim/core/stats"
)

type OnApplyEffects func(aura *Aura, sim *Simulation, target *Unit, spell *Spell)
type ApplySpellResults func(sim *Simulation, target *Unit, spell *Spell)
type ExpectedDamageCalculator func(sim *Simulation, target *Unit, spell *Spell, useSnapshot bool) *SpellResult
type CanCastCondition func(sim *Simulation, target *Unit) bool

type SpellConfig struct {
	// See definition of Spell (below) for comments on these.
	ActionID
	SpellSchool    SpellSchool
	ProcMask       ProcMask
	Flags          SpellFlag
	MissileSpeed   float64
	BaseCost       float64
	MetricSplits   int
	ClassSpellMask int64

	ManaCost   ManaCostOptions
	EnergyCost EnergyCostOptions
	RageCost   RageCostOptions
	RuneCost   RuneCostOptions
	FocusCost  FocusCostOptions

	Cast               CastConfig
	ExtraCastCondition CanCastCondition

	// Optional range constraints. If supplied, these are used to modify the ExtraCastCondition above to additionally check for DistanceFromTarget.
	MinRange float64
	MaxRange float64

	BonusHitPercent      float64
	BonusCritPercent     float64
	BonusSpellPower      float64
	BonusExpertiseRating float64

	DamageMultiplier         float64
	DamageMultiplierAdditive float64
	CritMultiplier           float64

	BonusCoefficient float64 // EffectBonusCoefficient in SpellEffect client DB table, "SP mod" on Wowhead (not necessarily shown there even if > 0)

	ThreatMultiplier float64

	FlatThreatBonus float64

	// Performs the actions of this spell.
	ApplyEffects ApplySpellResults

	// Optional field. Calculates expected average damage.
	ExpectedInitialDamage ExpectedDamageCalculator
	ExpectedTickDamage    ExpectedDamageCalculator

	Dot    DotConfig
	Hot    DotConfig
	Shield ShieldConfig

	RelatedAuras    []AuraArray
	RelatedDotSpell *Spell
}

type Spell struct {
	// ID for this spell.
	ActionID

	// The unit who will perform this spell.
	Unit *Unit

	// Fire, Frost, Shadow, etc.
	SpellSchool SpellSchool
	SchoolIndex stats.SchoolIndex

	// Controls which effects can proc from this spell.
	ProcMask ProcMask

	// Flags
	Flags SpellFlag

	// The specific class spell id
	// should be a unique bit
	ClassSpellMask int64

	// Speed in yards/second. Spell missile speeds can be found in the game data.
	// Example: https://wow.tools/dbc/?dbc=spellmisc&build=3.4.0.44996
	MissileSpeed float64

	ResourceMetrics *ResourceMetrics
	healthMetrics   []*ResourceMetrics

	Cost               SpellCost // Cost for the spell.
	DefaultCast        Cast      // Default cast parameters with all static effects applied.
	CD                 Cooldown
	SharedCD           Cooldown
	ExtraCastCondition CanCastCondition

	// Optional range constraints. If supplied, these are used to modify the ExtraCastCondition above to additionally check for DistanceFromTarget.
	MinRange float64
	MaxRange float64

	castTimeFn func(spell *Spell) time.Duration // allows to override CastTime()

	// Performs a cast of this spell.
	castFn CastSuccessFunc

	SpellMetrics      []SpellMetrics
	splitSpellMetrics [][]SpellMetrics // Used to split metrics by some condition.
	casts             int              // Sum of casts on all targets, for efficient CPM calculation

	// Performs the actions of this spell.
	ApplyEffects ApplySpellResults

	// Optional field. Calculates expected average damage.
	expectedInitialDamageInternal ExpectedDamageCalculator
	expectedTickDamageInternal    ExpectedDamageCalculator

	// The current or most recent cast data.
	CurCast Cast

	BonusHitPercent          float64
	BonusCritPercent         float64
	BonusSpellPower          float64
	BonusExpertiseRating     float64
	CastTimeMultiplier       float64
	CostMultiplier           float64
	CdMultiplier             float64
	DamageMultiplier         float64
	DamageMultiplierAdditive float64
	CritMultiplier           float64

	BonusCoefficient float64 // EffectBonusCoefficient in SpellEffect client DB table, "SP mod" on Wowhead (not necessarily shown there even if > 0)

	// Multiplier for all threat generated by this effect.
	ThreatMultiplier float64

	// Adds a fixed amount of threat to this spell, before multipliers.
	FlatThreatBonus float64

	resultCache SpellResult

	dots   DotArray
	aoeDot *Dot

	shields    ShieldArray
	selfShield *Shield

	// Per-target auras that are related to this spell, usually buffs or debuffs applied by the spell.
	RelatedAuras    []AuraArray
	RelatedDotSpell *Spell
}

func (unit *Unit) OnSpellRegistered(handler SpellRegisteredHandler) {
	for _, spell := range unit.Spellbook {
		handler(spell)
	}
	unit.spellRegistrationHandlers = append(unit.spellRegistrationHandlers, handler)
}

// Registers a new spell to the unit. Returns the newly created spell.
func (unit *Unit) RegisterSpell(config SpellConfig) *Spell {
	if len(unit.Spellbook) > 100 {
		panic(fmt.Sprintf("Over 100 registered spells when registering %s! There is probably a spell being registered every iteration.", config.ActionID))
	}

	// Default the other damage multiplier to 1 if only one or the other is set.
	if config.DamageMultiplier != 0 && config.DamageMultiplierAdditive == 0 {
		config.DamageMultiplierAdditive = 1
	} else if config.DamageMultiplierAdditive != 0 && config.DamageMultiplier == 0 {
		config.DamageMultiplier = 1
	}

	if (config.DamageMultiplier != 0 || config.ThreatMultiplier != 0) && config.ProcMask == ProcMaskUnknown {
		panic("ProcMask for spell " + config.ActionID.String() + " not set")
	}

	if (config.DamageMultiplier != 0 || config.ThreatMultiplier != 0) && config.SpellSchool == SpellSchoolNone {
		panic("SpellSchool for spell " + config.ActionID.String() + " not set")
	}

	if config.Cast.CD.Timer != nil && config.Cast.CD.Duration == 0 {
		panic("Cast.CD w/o Duration specified for spell " + config.ActionID.String())
	}

	if config.Cast.SharedCD.Timer != nil && config.Cast.SharedCD.Duration == 0 {
		panic("Cast.SharedCD w/o Duration specified for spell " + config.ActionID.String())
	}

	if config.Cast.CastTime == nil {
		config.Cast.CastTime = func(spell *Spell) time.Duration {
			return spell.Unit.ApplyCastSpeedForSpell(spell.DefaultCast.CastTime, spell)
		}
	}

	spell := &Spell{
		ActionID:       config.ActionID,
		Unit:           unit,
		SpellSchool:    config.SpellSchool,
		ProcMask:       config.ProcMask,
		Flags:          config.Flags,
		MissileSpeed:   config.MissileSpeed,
		ClassSpellMask: config.ClassSpellMask,

		DefaultCast:        config.Cast.DefaultCast,
		CD:                 config.Cast.CD,
		SharedCD:           config.Cast.SharedCD,
		ExtraCastCondition: config.ExtraCastCondition,

		castTimeFn: config.Cast.CastTime,

		ApplyEffects: config.ApplyEffects,

		expectedInitialDamageInternal: config.ExpectedInitialDamage,
		expectedTickDamageInternal:    config.ExpectedTickDamage,

		BonusHitPercent:          config.BonusHitPercent,
		BonusCritPercent:         config.BonusCritPercent,
		BonusSpellPower:          config.BonusSpellPower,
		BonusExpertiseRating:     config.BonusExpertiseRating,
		CastTimeMultiplier:       1,
		CostMultiplier:           1,
		CdMultiplier:             1,
		DamageMultiplier:         config.DamageMultiplier,
		DamageMultiplierAdditive: config.DamageMultiplierAdditive,
		CritMultiplier:           config.CritMultiplier,

		BonusCoefficient: config.BonusCoefficient,

		ThreatMultiplier: config.ThreatMultiplier,
		FlatThreatBonus:  config.FlatThreatBonus,

		splitSpellMetrics: make([][]SpellMetrics, max(1, config.MetricSplits)),

		RelatedAuras:    config.RelatedAuras,
		RelatedDotSpell: config.RelatedDotSpell,
	}

	switch {
	case spell.SpellSchool.Matches(SpellSchoolPhysical):
		spell.SchoolIndex = stats.SchoolIndexPhysical
	case spell.SpellSchool.Matches(SpellSchoolArcane):
		spell.SchoolIndex = stats.SchoolIndexArcane
	case spell.SpellSchool.Matches(SpellSchoolFire):
		spell.SchoolIndex = stats.SchoolIndexFire
	case spell.SpellSchool.Matches(SpellSchoolFrost):
		spell.SchoolIndex = stats.SchoolIndexFrost
	case spell.SpellSchool.Matches(SpellSchoolHoly):
		spell.SchoolIndex = stats.SchoolIndexHoly
	case spell.SpellSchool.Matches(SpellSchoolNature):
		spell.SchoolIndex = stats.SchoolIndexNature
	case spell.SpellSchool.Matches(SpellSchoolShadow):
		spell.SchoolIndex = stats.SchoolIndexShadow
	}

	// newXXXCost() all update spell.DefaultCast.Cost
	if config.ManaCost.BaseCost != 0 || config.ManaCost.FlatCost != 0 {
		spell.Cost = newManaCost(spell, config.ManaCost)
	} else if config.EnergyCost.Cost != 0 {
		spell.Cost = newEnergyCost(spell, config.EnergyCost)
	} else if config.RageCost.Cost != 0 {
		spell.Cost = newRageCost(spell, config.RageCost)
	} else if config.RuneCost.BloodRuneCost != 0 || config.RuneCost.FrostRuneCost != 0 || config.RuneCost.UnholyRuneCost != 0 || config.RuneCost.RunicPowerCost != 0 || config.RuneCost.RunicPowerGain != 0 {
		spell.Cost = newRuneCost(spell, config.RuneCost)
	} else if config.FocusCost.Cost != 0 {
		spell.Cost = newFocusCost(spell, config.FocusCost)
	}

	spell.createDots(config.Dot, false)
	spell.createDots(config.Hot, true)
	spell.createShields(config.Shield)

	var emptyCast Cast

	if spell.DefaultCast == emptyCast && spell.Cost != nil {
		panic("Empty DefaultCast with a cost for spell " + config.ActionID.String())
	}

	if spell.DefaultCast.GCD <= 0 && spell.DefaultCast.CastTime == 0 {
		config.Cast.IgnoreHaste = true
	}

	if spell.DefaultCast == emptyCast {
		if config.ExtraCastCondition == nil && config.Cast.CD.Timer == nil && config.Cast.SharedCD.Timer == nil {
			spell.castFn = spell.makeCastFuncAutosOrProcs()
		} else {
			spell.castFn = spell.makeCastFuncSimple()
		}
	} else {
		spell.castFn = spell.makeCastFunc(config.Cast)
	}

	if spell.ApplyEffects == nil {
		spell.ApplyEffects = func(*Simulation, *Unit, *Spell) {}
	}

	// Apply range constraints if requested. This is done after generating the castFn
	// for performance reasons, so that auto-attacks can be managed separately during
	// movement actions rather than constantly polling range checks.
	if (config.MinRange != 0) || (config.MaxRange != 0) {
		spell.MinRange = config.MinRange
		spell.MaxRange = config.MaxRange
		oldExtraCastCondition := spell.ExtraCastCondition
		spell.ExtraCastCondition = func(sim *Simulation, target *Unit) bool {
			if ((spell.MinRange != 0) && (spell.Unit.DistanceFromTarget < spell.MinRange)) || ((spell.MaxRange != 0) && (spell.Unit.DistanceFromTarget > spell.MaxRange)) {
				if sim.Log != nil {
					sim.Log("Cannot cast spell %s, out of range!", spell.ActionID)
				}

				return false
			}

			return (oldExtraCastCondition == nil) || oldExtraCastCondition(sim, target)
		}
	}

	unit.Spellbook = append(unit.Spellbook, spell)

	for _, handler := range unit.spellRegistrationHandlers {
		handler(spell)
	}

	if unit.Env != nil && unit.Env.IsFinalized() {
		spell.finalize()
	}

	return spell
}

// Returns the first registered spell with the given ID, or nil if there are none.
func (unit *Unit) GetSpell(actionID ActionID) *Spell {
	for _, spell := range unit.Spellbook {
		if spell.ActionID.SameAction(actionID) {
			return spell
		}
	}
	return nil
}

// Retrieves an existing spell with the same ID as the config uses, or registers it if there is none.
func (unit *Unit) GetOrRegisterSpell(config SpellConfig) *Spell {
	registered := unit.GetSpell(config.ActionID)
	if registered == nil {
		return unit.RegisterSpell(config)
	} else {
		return registered
	}
}

func (spell *Spell) Dot(target *Unit) *Dot {
	if spell.dots == nil {
		if spell.RelatedDotSpell != nil {
			return spell.RelatedDotSpell.Dot(target)
		}
		return nil
	}
	return spell.dots.Get(target)
}
func (spell *Spell) CurDot() *Dot {
	if spell.dots == nil {
		if spell.RelatedDotSpell != nil {
			return spell.RelatedDotSpell.CurDot()
		}
		return nil
	}
	return spell.dots.Get(spell.Unit.CurrentTarget)
}
func (spell *Spell) AOEDot() *Dot {
	return spell.aoeDot
}
func (spell *Spell) Hot(target *Unit) *Dot {
	if spell.dots == nil {
		return nil
	}
	return spell.dots.Get(target)
}
func (spell *Spell) CurHot() *Dot {
	if spell.dots == nil {
		return nil
	}
	return spell.dots.Get(spell.Unit.CurrentTarget)
}
func (spell *Spell) AOEHot() *Dot {
	return spell.aoeDot
}
func (spell *Spell) SelfHot() *Dot {
	return spell.aoeDot
}
func (spell *Spell) Shield(target *Unit) *Shield {
	if spell.shields == nil {
		return nil
	}
	return spell.shields.Get(target)
}
func (spell *Spell) SelfShield() *Shield {
	return spell.selfShield
}

// Metrics for the current iteration
func (spell *Spell) CurDamagePerCast() float64 {
	if spell.SpellMetrics[0].Casts == 0 {
		return 0
	} else {
		casts := int32(0)
		damage := 0.0
		for _, opponent := range spell.Unit.GetOpponents() {
			casts += spell.SpellMetrics[opponent.UnitIndex].Casts
			damage += spell.SpellMetrics[opponent.UnitIndex].TotalDamage
		}
		return damage / float64(casts)
	}
}

// Current casts per minute
func (spell *Spell) CurCPM(sim *Simulation) float64 {
	if sim.CurrentTime <= 0 {
		return 0
	}
	casts := float64(spell.casts)
	minutes := float64(sim.CurrentTime) / float64(time.Minute)
	return casts / minutes
}

func (spell *Spell) finalize() {
	if len(spell.splitSpellMetrics) > 1 && spell.ActionID.Tag != 0 {
		panic(spell.ActionID.String() + " has split metrics and a non-zero tag, can only have one!")
	}
	for i := range spell.splitSpellMetrics {
		spell.splitSpellMetrics[i] = make([]SpellMetrics, len(spell.Unit.Env.AllUnits))
	}
	spell.SpellMetrics = spell.splitSpellMetrics[0]
}

func (spell *Spell) reset(_ *Simulation) {
	for i := range spell.splitSpellMetrics {
		for j := range spell.SpellMetrics {
			spell.splitSpellMetrics[i][j] = SpellMetrics{}
		}
	}
	spell.casts = 0
}

func (spell *Spell) SetMetricsSplit(splitIdx int32) {
	spell.SpellMetrics = spell.splitSpellMetrics[splitIdx]
	spell.ActionID.Tag = splitIdx
}

func (spell *Spell) doneIteration() {
	if spell.Flags.Matches(SpellFlagNoMetrics) {
		return
	}

	if len(spell.splitSpellMetrics) == 1 {
		spell.Unit.Metrics.addSpellMetrics(spell, spell.ActionID, spell.SpellMetrics)
	} else {
		for i, spellMetrics := range spell.splitSpellMetrics {
			spell.Unit.Metrics.addSpellMetrics(spell, spell.ActionID.WithTag(int32(i)), spellMetrics)
		}
	}
}

func (spell *Spell) HealthMetrics(target *Unit) *ResourceMetrics {
	if spell.healthMetrics == nil {
		spell.healthMetrics = make([]*ResourceMetrics, len(spell.Unit.AttackTables))
	}
	if spell.healthMetrics[target.UnitIndex] == nil {
		spell.healthMetrics[target.UnitIndex] = target.NewHealthMetrics(spell.ActionID)
	}
	return spell.healthMetrics[target.UnitIndex]
}

func (spell *Spell) ReadyAt() time.Duration {
	return BothTimersReadyAt(spell.CD.Timer, spell.SharedCD.Timer)
}

func (spell *Spell) IsReady(sim *Simulation) bool {
	if spell == nil {
		return false
	}
	return BothTimersReady(spell.CD.Timer, spell.SharedCD.Timer, sim)
}

func (spell *Spell) TimeToReady(sim *Simulation) time.Duration {
	return MaxTimeToReady(spell.CD.Timer, spell.SharedCD.Timer, sim)
}

// Returns whether a call to Cast() would be successful, without actually
// doing a cast.
func (spell *Spell) CanCast(sim *Simulation, target *Unit) bool {
	if spell == nil {
		return false
	}

	if spell.ExtraCastCondition != nil && !spell.ExtraCastCondition(sim, target) {
		//if sim.Log != nil {
		//	sim.Log("Cant cast because of extra condition")
		//}
		return false
	}

	// While moving only instant casts are possible
	if spell.Flags&SpellFlagCanCastWhileMoving == 0 && spell.DefaultCast.CastTime > 0 && spell.Unit.Moving {
		//if sim.Log != nil {
		//	sim.Log("Cant cast because moving")
		//}
		return false
	}

	// While casting or channeling, no other action is possible
	if spell.Unit.Hardcast.Expires > sim.CurrentTime {
		//if sim.Log != nil {
		//	sim.Log("Cant cast because already casting/channeling")
		//}
		return false
	}

	if ((spell.DefaultCast.GCD > 0) || (spell.Flags.Matches(SpellFlagMCD) && spell.Unit.Rotation.inSequence)) && !spell.Unit.GCD.IsReady(sim) {
		//if sim.Log != nil {
		//	sim.Log("Cant cast because of GCD")
		//}
		return false
	}

	if !BothTimersReady(spell.CD.Timer, spell.SharedCD.Timer, sim) {
		//if sim.Log != nil {
		//	sim.Log("Cant cast because of CDs")
		//}
		return false
	}

	if spell.Cost != nil {
		// temp hack
		spell.CurCast.Cost = spell.DefaultCast.Cost
		if !spell.Cost.MeetsRequirement(sim, spell) {
			//if sim.Log != nil {
			//	sim.Log("Cant cast because of resource cost")
			//}
			return false
		}
	}

	return true
}

func (spell *Spell) Cast(sim *Simulation, target *Unit) bool {
	if spell.DefaultCast.EffectiveTime() > 0 {
		spell.Unit.CancelQueuedSpell(sim)
	}
	if target == nil {
		target = spell.Unit.CurrentTarget
	}
	return spell.castFn(sim, target)
}

// Skips the actual cast and applies spell effects immediately.
func (spell *Spell) SkipCastAndApplyEffects(sim *Simulation, target *Unit) {
	if sim.Log != nil && !spell.Flags.Matches(SpellFlagNoLogs) {
		spell.Unit.Log(sim, "Casting %s (Cost = %0.03f, Cast Time = %s)",
			spell.ActionID, spell.DefaultCast.Cost, time.Duration(0))
		spell.Unit.Log(sim, "Completed cast %s", spell.ActionID)
	}
	spell.applyEffects(sim, target)
}

func (spell *Spell) applyEffects(sim *Simulation, target *Unit) {
	spell.SpellMetrics[target.UnitIndex].Casts++
	spell.casts++

	// Not sure if we want to split this flag into its own?
	// Both are used to optimize away unneccesery calls and 99%
	// of the time are gonna be used together. For now just in one
	if !spell.Flags.Matches(SpellFlagNoOnCastComplete) {
		spell.Unit.OnApplyEffects(sim, target, spell)
	}

	spell.ApplyEffects(sim, target, spell)
}

func (spell *Spell) ApplyAOEThreatIgnoreMultipliers(threatAmount float64) {
	numTargets := spell.Unit.Env.GetNumTargets()
	for i := int32(0); i < numTargets; i++ {
		spell.SpellMetrics[i].TotalThreat += threatAmount
	}
}
func (spell *Spell) ApplyAOEThreat(threatAmount float64) {
	spell.ApplyAOEThreatIgnoreMultipliers(threatAmount * spell.Unit.PseudoStats.ThreatMultiplier)
}

func (spell *Spell) finalizeExpectedDamage(result *SpellResult) {
	if !spell.SpellSchool.Matches(SpellSchoolPhysical) {
		result.Damage /= result.ResistanceMultiplier
		result.Damage *= AverageMagicPartialResistMultiplier
		result.ResistanceMultiplier = AverageMagicPartialResistMultiplier
	}
	result.inUse = false
}
func (spell *Spell) ExpectedInitialDamage(sim *Simulation, target *Unit) float64 {
	result := spell.expectedInitialDamageInternal(sim, target, spell, false)
	spell.finalizeExpectedDamage(result)
	return result.Damage
}
func (spell *Spell) ExpectedTickDamage(sim *Simulation, target *Unit) float64 {
	result := spell.expectedTickDamageInternal(sim, target, spell, false)
	spell.finalizeExpectedDamage(result)
	return result.Damage
}
func (spell *Spell) ExpectedTickDamageFromCurrentSnapshot(sim *Simulation, target *Unit) float64 {
	result := spell.expectedTickDamageInternal(sim, target, spell, true)
	spell.finalizeExpectedDamage(result)
	return result.Damage
}

// Time until either the cast is finished or GCD is ready again, whichever is longer
func (spell *Spell) EffectiveCastTime() time.Duration {
	// TODO: this is wrong for spells like shadowfury, that have a GCD of less than 1s
	return max(spell.Unit.SpellGCD(),
		spell.Unit.ApplyCastSpeedForSpell(spell.DefaultCast.EffectiveTime(), spell))
}

// Time until the cast is finished (ignoring GCD)
func (spell *Spell) CastTime() time.Duration {
	return spell.castTimeFn(spell)
}

func (spell *Spell) TravelTime() time.Duration {
	if spell.MissileSpeed == 0 {
		return 0
	} else {
		return time.Duration(float64(time.Second) * spell.Unit.DistanceFromTarget / spell.MissileSpeed)
	}
}

// Returns true if the given mask matches the spell mask
func (spell *Spell) Matches(mask int64) bool {
	return spell.ClassSpellMask&mask > 0
}

// Handles computing the cost of spells and checking whether the Unit
// meets them.
type SpellCost interface {
	// Whether the Unit associated with the spell meets the resource cost
	// requirements to cast the spell.
	MeetsRequirement(*Simulation, *Spell) bool

	// Returns a message for when the cast fails due to lack of resources.
	CostFailureReason(*Simulation, *Spell) string

	// Subtracts the resources used from a cast from the Unit.
	SpendCost(*Simulation, *Spell)

	// Space for handling refund mechanics. Not all spells provide refunds.
	IssueRefund(*Simulation, *Spell)
}

func (spell *Spell) IssueRefund(sim *Simulation) {
	spell.Cost.IssueRefund(sim, spell)
}
