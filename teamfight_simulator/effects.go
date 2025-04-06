package main

type Effect struct {
	Name     EffectName
	Identity int

	Target Unit // Affected unit
	Source Unit

	Start    int
	Duration int
	Amount   float64

	// overrideStackable StackOverride
}

type EffectName int

const (
	AllEffects EffectName = iota
	Sunder                // Reduce Resistance
	Shred
	Burn  // Deal True Damage Every Tick
	Wound // Reduce healing
	Stun
	Held // Cannot be removed like a stun unless the holder gets disabled
	Shield
	Chill
	Invulnerability
	Untargetability
)

func IsStackable(name EffectName) bool {
	switch name {
	case Shield:
		return true
	default:
		return false
	}
}

type EffectDescription struct {
	name     EffectName
	duration int
	amount   float64

	target Unit
	source Unit
}

func EffectBuilder(name EffectName) *EffectDescription {
	return &EffectDescription{
		name: name,
	}
}

// Also sets Target as the Source of the effect is Source is not yet set
func (description *EffectDescription) Target(unit Unit) *EffectDescription {
	description.target = unit
	if description.source == nil {
		description.source = unit
	}
	return description
}

func (description *EffectDescription) Source(unit Unit) *EffectDescription {
	description.source = unit
	return description
}

func (description *EffectDescription) Duration(duration int) *EffectDescription {
	description.duration = duration
	return description
}

func (description *EffectDescription) Amount(amount float64) *EffectDescription {
	description.amount = amount
	return description
}

func ContainsEffect(unitEffects map[EffectName]bool, effect EffectName) bool {
	_, ok := unitEffects[effect]
	return ok
}

func EffectRemaining(effect *Effect, combat *Combat) int {
	// LogC(combat, fmt.Sprintf("Effect%s remaining %d", string(effect.Identity), effect.Duration+effect.Start-combat.Tick))
	return effect.Duration + effect.Start - combat.CurrentTick
}

func GetActiveEffects(unit Unit, effectType EffectName) []*Effect {
	// TODO: Any inactive effects will be pruned
	out := make([]*Effect, 0, len(unit.Effects()))
	for _, effect := range unit.Effects() {
		isSelected := effectType == AllEffects || effectType == effect.Name
		if isSelected && EffectRemaining(effect, unit.Combat()) > 0 {
			out = append(out, effect)
		}
	}

	return out
}

func GetActiveEffectPresenceMap(unit Unit) map[EffectName]bool {
	out := make(map[EffectName]bool)
	for _, effect := range unit.Effects() {
		if EffectRemaining(effect, unit.Combat()) > 0 {
			out[effect.Name] = true
		}
	}
	return out
}

func GetEffectLongestRemainingDuration(unit Unit, effectType EffectName) int {
	effects := GetActiveEffects(unit, effectType)
	duration := -1

	for _, effect := range effects {
		if EffectRemaining(effect, unit.Combat()) > duration {
			duration = EffectRemaining(effect, unit.Combat())
		}
	}
	return duration
}

func GetActiveEffectSum(unit Unit, effectType EffectName) (float64, Unit) {
	effects := GetActiveEffects(unit, effectType)

	if IsStackable(effectType) {
		sum := 0.0
		for _, effect := range effects {
			sum += effect.Amount
		}

		// LogC(unit.Combat, fmt.Sprintf("len:%d, Effect%d sum: %f\n", len(effects), int(effectType), sum))
		return sum, nil
	}

	maximum := 0.0
	var author Unit
	for _, effect := range effects {
		if effect.Amount > maximum {
			maximum = effect.Amount
		}
	}
	return maximum, author
}

// TODO: Unit here
func DeactivateEffectEarly(combat *Combat, effect *Effect) {
	if EffectRemaining(effect, combat) <= 0 {
		return
	}
	effect.Duration = combat.CurrentTick - effect.Start
}
