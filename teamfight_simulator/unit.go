package main

import (
	"math"
	"slices"
)

const defaultCriticalChance = .25
const defaultCriticalDamage = 1.4
const defaultAbilityPower = 100.0
const defaultDamageAmp = 0.0
const defaultDurability = 0.0
const defaultOmnivamp = 0.0
const MaximumChaseDistance = 1
const DefaultStartMana = 0

type TeamId int

type TraitId string

type Unit interface {
	Vitals

	/************ [Boolean] ***********/
	IsBackStarter() bool
	IsFrontStarter() bool
	IsAlive() bool

	AddModifier(modifier *Modifier) func()

	// Effects
	AddEffect(effect *Effect)
	ContainsEffect(effect EffectName) bool
	DeleteEffect(effect *Effect)
	Effects() []*Effect

	/*********** [Setters] ***********/
	PrepareForNextCombat()

	AddHealth(amount int)
	AddMana(amount int)
	AddItem(item ItemName)
	AddDamager(unit Unit)

	SetHealth(amount int)
	SetPosition(position int)
	SetTarget(target Unit)

	ResetDamagers()
	ResetChaseDistance()

	IncrementChaseDistance()

	/************ [Events] ***********/
	Notify(event *Event)
	RegisterEvent(eventType EventType, handler func(*Event))

	/************ [Board Calculations] ***********/
	SearchUnits(options *SearchOptions) []Unit

	/************ [State] ***********/
	SetState(state UnitStateName, duration int)
	StateName() UnitStateName
	StateRemaining() int
	CombatReport() *CombatReport

	/************ [Combat] ***********/
	DealEffect(options *EffectDescription) *Effect

	/************ [Getters] ***********/
	Name() string
	Id() int
	Combat() *Combat
	Position() int
	Team() TeamId
	Level() int
	Target() Unit
	CurrentHealth() int
	MissingHealth() int
	CurrentMana() int
	Damagers() []Unit
	ChaseDistance() int
	CountTraits(trait TraitId) map[TraitId]int
	Items() []ItemName
}

type UnitData struct {
	Prototype *Champion
	combat    *Combat

	addModif  map[ModifierName][]Modifier
	multModif map[ModifierName][]Modifier

	id       int
	position int
	level    int
	team     TeamId

	items []ItemName

	state UnitState

	effects map[int]*Effect
	report  *CombatReport

	damagers []Unit

	currentHealth int
	currentMana   int

	target Unit

	// TODO: Add Property that allows us to specify that Target Dummies cannot move!
	// (Since target dummies will be a unit by themselves)
	// Perhaps we can override process tick for each unit
	eventNotifiers map[EventType][]func(*Event)
	interruptors   map[int][]func()

	processActiveAbility func(Unit, func())
	chaseDistance        int
}

type UnitStateName int

const (
	IdleState = iota
	MovingState
	AttackAnimationState
	CastAnimationState
)

func (unit *UnitData) Combat() *Combat {
	return unit.combat
}

func (unit *UnitData) Id() int {
	return unit.id
}

func (unit *UnitData) SetPosition(position int) {
	unit.position = position
}

func (unit *UnitData) AddModifier(modifier *Modifier) func() {
	panic("Unimplemented: return modifier id")
}

func (unit *UnitData) DealEffect(options *EffectDescription) *Effect {
	return unit.Combat().DealEffect(options)
}

func (unit *UnitData) Damagers() []Unit {
	return unit.damagers
}

func (unit *UnitData) AddDamager(damager Unit) {
	unit.damagers = append(unit.damagers, damager)
}

func (unit *UnitData) ResetDamagers() {
	unit.damagers = make([]Unit, 0, len(unit.damagers))
}

func (unit *UnitData) ChaseDistance() int {
	return unit.chaseDistance
}

func (unit *UnitData) IncrementChaseDistance() {
	unit.chaseDistance++
}

func (unit *UnitData) ResetChaseDistance() {
	unit.chaseDistance = 0
}

func (unit *UnitData) AddItem(item ItemName) {
	// TODO: Register item here as well????
	unit.items = append(unit.items, item)
}

func (unit *UnitData) SetTarget(target Unit) {
	unit.target = target
}

func (unit *UnitData) SetHealth(amount int) {
	unit.currentHealth = amount
}

func (unit *UnitData) Target() Unit {
	return unit.target
}

func (unit *UnitData) Items() []ItemName {
	return unit.items
}

func (unit *UnitData) RemoveItem(item ItemName) {
	panic("Unimplemented")
}

func (unit *UnitData) CountTraits(TraitId) map[TraitId]int {
	panic("Unimplemented")
}

func (unit *UnitData) ContainsEffect(effectType EffectName) bool {
	panic("Unimplemented!")
}

func (unit *UnitData) StateRemaining() int {
	return unit.state.Start + unit.state.Duration - unit.combat.CurrentTick
}

func (unit *UnitData) StateName() UnitStateName {
	return unit.state.Name
}

func (unit *UnitData) SetState(newState UnitStateName, duration int) {
	unit.state.Name = newState
	unit.state.Start = unit.combat.CurrentTick
	unit.state.Duration = duration
}

type UnitState struct {
	Name     UnitStateName
	Start    int
	Duration int //TODO: some casts have an arbitrary duration, ex: caitlyn on retarget
}

func (unit *UnitData) AddHealth(amount int) {
	unit.currentHealth += amount
	// TODO: We may change this implementation in the future
	// different from heal as it does not get affected by wound
}

func (unit *UnitData) Level() int {
	return unit.level
}

func (unit *UnitData) Team() TeamId {
	return unit.team
}

func RoundStat(stat float64) int {
	return int(math.Round(stat))
}

func (unit *UnitData) RegisterEvent(eventType EventType, handler func(*Event)) {
	panic("Unimplemented")
}

func (unit *UnitData) Notify(event *Event) {
	panic("Unimplemented")
}

func (unit *UnitData) Position() int {
	return unit.position
}

func (unit *UnitData) UnitRegisterEvent(eventType EventType, handler func(*Event)) {
	unit.eventNotifiers[EventType(eventType)] = append(unit.eventNotifiers[EventType(eventType)], handler)

}

func (unit *UnitData) UnitNotifyEvent(e *Event) {
	functions, ok := unit.eventNotifiers[e.Type]

	if ok {
		for _, foo := range functions {
			foo(e)
		}
	}
}

func (unit *UnitData) AddInterruptibleAction(cancel func(), expiration int) {

}

func (unit *UnitData) Interrupt() {
	for _, cancels := range unit.interruptors {
		for _, cancel := range cancels {
			cancel()
		}
	}
	clear(unit.interruptors)
}

func (unit *UnitData) CurrentHealth() int {
	return unit.currentHealth
}

func (unit *UnitData) CurrentMana() int {
	return unit.currentMana
}

func (unit *UnitData) MaxMana() int {
	return int(math.Round(unit.Prototype.Mana))
}

func (unit *UnitData) Name() string {
	return unit.Prototype.Name
}

func (unit *UnitData) IsAlive() bool {
	return unit.currentHealth >= 0
}

func (unit *UnitData) SearchUnits(options *SearchOptions) []Unit {
	units := unit.Combat().Units
	out := make([]Unit, 0, len(units))

	for _, otherUnit := range units {
		valid := true
		for _, filter := range options.filters {
			if !filter(otherUnit) {
				valid = false
				break
			}
		}

		if valid {
			out = append(out, otherUnit)
		}
	}

	slices.SortFunc(out, options.cmp)
	upper := MinInt(options.limit, len(out))
	return out[:upper]
}

func (unit *UnitData) IsFrontStarter() bool {
	row := unit.position / 7
	if unit.team == 0 {
		return !(row == 6 || row == 7)
	}
	return !(row == 0 || row == 1)
}

func (unit *UnitData) IsBackStarter() bool {
	row := unit.position / 7
	if unit.team == 0 {
		return row == 6 || row == 7
	}
	return row == 0 || row == 1
}

func (unit *UnitData) PrepareForNextCombat() {

}

func (unit *UnitData) AddMana(count int) {
	unit.currentMana += count
}

func (unit *UnitData) CombatReport() *CombatReport {
	return unit.report
}

func (unit *UnitData) AddEffect(effect *Effect) {
	unit.effects[effect.Identity] = effect
}

func (unit *UnitData) DeleteEffect(effect *Effect) {
	delete(unit.effects, effect.Identity)

}
func (unit *UnitData) Effects() []*Effect {
	out := make([]*Effect, len(unit.effects))
	for _, effect := range unit.effects {
		out = append(out, effect)
	}
	return out
}

const ALL = 0

const NeverExpires = -1

type ModifierType int

const (
	Additive       ModifierType = 0
	Multiplicative ModifierType = 1
)

type Modifier struct {
	Name       ModifierName
	Type       ModifierType
	Data       map[string]float64
	Retriever  func(Unit) float64
	Expiration int
}

func AdditiveModifier(name ModifierName) *Modifier {
	return &Modifier{
		Name:       name,
		Type:       Additive,
		Expiration: NeverExpires,
	}
}

func MultiplicativeModifier(name ModifierName) *Modifier {
	return &Modifier{
		Name:       name,
		Type:       Multiplicative,
		Expiration: NeverExpires,
	}
}

func (modifier *Modifier) WithDynamicData(data map[string]float64) *Modifier {
	modifier.Data = data
	return modifier
}

func (modifier *Modifier) WithRetriever(retriever func(Unit) float64) *Modifier {
	modifier.Retriever = retriever
	return modifier
}

func (modifier *Modifier) WithConstant(constant float64) *Modifier {
	modifier.Retriever = func(u Unit) float64 { return constant }
	return modifier
}

func (modifier *Modifier) WithExpiration(expiration int) *Modifier {
	modifier.Expiration = expiration
	return modifier
}

type ModifierName string

const (
	AttackDamage         ModifierName = "AttackDamage"
	AttackSpeed          ModifierName = "AttackSpeed"
	AbilityPower         ModifierName = "AbilityPower"
	StartMana            ModifierName = "StartMana"
	MaximumMana          ModifierName = "MaximumMana"
	ArmorResist          ModifierName = "ArmorResist"
	MagicResist          ModifierName = "MagicResist"
	MaximumHealth        ModifierName = "MaximumHealth"
	CriticalStrikeChance ModifierName = "CriticalStrikeChance"
	CriticalStrikeDamage ModifierName = "CriticalStrikeDamage"
	Omnivamp             ModifierName = "Omnivamp"
	Durabiltiy           ModifierName = "Durability"
	DamageAmp            ModifierName = "DamageAmp" // TODO: Does damage amp amplify True Damage?
	HexRange             ModifierName = "HexRange"
)

type Vitals interface {
	// Group 1: Vitals
	StartMana() int
	MaxMana() int
	MaxHealth() int

	// Group 2: Damage per second
	AttackDamage() int
	AttackSpeed() float64
	CriticalStrikeChance() float64
	CriticalStrikeDamage() float64

	// Group 3: Ability Power
	AbilityPower() int

	// Group 4: Resisters
	ArmorResist() int
	MagicResist() int

	// Group 5: Ampers
	DamageAmp() float64
	Durability() float64
	Omnivamp() float64

	// Group 6: Other
	HexRange() int
}

func sumMultiplicativeModifiers(unit *UnitData, name ModifierName) float64 {
	out := 1.0
	for _, modifier := range unit.multModif[name] {
		out += modifier.Retriever(Unit(unit))
	}
	return out
}

func sumAdditiveModifiers(unit *UnitData, name ModifierName) float64 {
	out := 0.0
	for _, modifier := range unit.addModif[name] {
		out += modifier.Retriever(Unit(unit))
	}
	return out
}

func sumApplyRoundModifiers(unit *UnitData, name ModifierName, baseStat float64) int {
	baseStat += sumAdditiveModifiers(unit, name)
	return RoundStat(baseStat * sumMultiplicativeModifiers(unit, name))
}

func sumApplyFloatModifiers(unit *UnitData, name ModifierName, baseStat float64) float64 {
	baseStat += sumAdditiveModifiers(unit, name)
	return baseStat * sumMultiplicativeModifiers(unit, name)
}

func (unit *UnitData) MaxHealth() int {
	protoHealth := unit.Prototype.Health[unit.level]
	return sumApplyRoundModifiers(unit, MaximumHealth, protoHealth)
}

func (unit *UnitData) AttackDamage() int {
	baseAttackDamage := unit.Prototype.AttackDamage[unit.level]
	return sumApplyRoundModifiers(unit, AttackDamage, baseAttackDamage)
}

func (unit *UnitData) AbilityPower() int {
	return sumApplyRoundModifiers(unit, AbilityPower, defaultAbilityPower)
}

func (unit *UnitData) ArmorResist() int {
	return sumApplyRoundModifiers(unit, ArmorResist, unit.Prototype.Armor)
}

func (unit *UnitData) MagicResist() int {
	return sumApplyRoundModifiers(unit, MagicResist, unit.Prototype.MagicResist)
}

func (unit *UnitData) AttackSpeed() float64 {
	// TODO: Make sure the chill calculation works as expected
	chill, _ := GetActiveEffectSum(unit, Chill)
	attackSpeed := sumApplyFloatModifiers(unit, AttackSpeed, unit.Prototype.AttackSpeed)
	return attackSpeed - attackSpeed*chill
}

func (unit *UnitData) StartMana() int {
	return RoundStat(sumApplyFloatModifiers(unit, StartMana, DefaultStartMana))

}

func (unit *UnitData) DamageAmp() float64 {
	return sumApplyFloatModifiers(unit, DamageAmp, defaultDamageAmp)
}

func (unit *UnitData) Durability() float64 {
	return sumApplyFloatModifiers(unit, Durabiltiy, defaultDurability)
}

func (unit *UnitData) Omnivamp() float64 {
	return sumApplyFloatModifiers(unit, Omnivamp, defaultOmnivamp)
}

func (unit *UnitData) CriticalStrikeChance() float64 {
	// TODO: Ensure IE and JG do not double dip into the 10% bonus
	return sumApplyFloatModifiers(unit, CriticalStrikeChance, defaultCriticalChance)
}

func (unit *UnitData) CriticalStrikeDamage() float64 {
	return sumApplyFloatModifiers(unit, CriticalStrikeDamage, defaultCriticalDamage)
}

func (unit *UnitData) HexRange() int {
	return sumApplyRoundModifiers(unit, HexRange, float64(unit.Prototype.HexRange))
}

func (unit *UnitData) MissingHealth() int {
	protoHealth := unit.Prototype.Health[unit.level]
	maxHealth := sumApplyRoundModifiers(unit, MaximumHealth, protoHealth)
	return maxHealth - unit.currentHealth
}
