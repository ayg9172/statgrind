package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"sort"
	"time"
)

// TODO: Always check if target is in range because target may move and champion can move too

// TODO: Melee champions not only have to be adjacent to target
// TODO: Melee champions must be done moving and their target must be done moving
const DefaultTickRate = 30.33

const LogLength = 30

type DamageType int

const (
	AttackDamageType DamageType = iota
	AbilityDamageType
	TrueDamageType
)

type Combat struct {
	PlayerOne *Player
	PlayerTwo *Player

	Units         []Unit
	Board         []Unit
	UnitRetriever map[int]Unit
	CombatLog     []string

	Portal      string
	Stage       int
	CurrentTick int
	WinningTeam TeamId

	NextId func() int

	TargetSelector     func(int, *Combat, TeamId) ([]Unit, []int)
	GetNeighbors       func(int) []int
	GetLogicalDistance func(int, int) int

	scheduling *Scheduling

	prototypes map[string]*Champion
}

type CombatReport struct {
	AttackDamage  int
	AbilityDamage int
	TrueDamage    int

	AttackResist  int
	AbilityResist int
	ShieldsResist int

	ShieldsBlock int
	HealthBlock  int
	Healing      int

	SunderDealt int
	ShredDealt  int
	WoundDealt  int
}

func CreateDefaultCombat() *Combat {
	out := &Combat{
		NextId:             CreateIdFactory(),
		PlayerOne:          CreateDefaultPlayer(),
		PlayerTwo:          CreateDefaultPlayer(),
		Units:              make([]Unit, 0, 20),
		UnitRetriever:      make(map[int]Unit),
		Board:              make([]Unit, 8*7),
		Portal:             "",
		Stage:              1,
		CurrentTick:        0,
		GetNeighbors:       CreateRandomGetNeighborsFunc(),
		TargetSelector:     CreateTargetSelector(),
		GetLogicalDistance: CreateGetLogicalDistanceFunction8x7(),
		WinningTeam:        -1,
	}
	return out
}

// Clears combat squares but retains the precomputed values
// TODO: Implement
func ClearCombat(combat *Combat) {

}

/****************** Other functions ******************/

func (combat *Combat) NewUnit(name string, level int, position int, team TeamId) Unit {
	prototype := combat.prototypes[name]
	unit := &UnitData{
		combat:         combat,
		id:             combat.NextId(),
		Prototype:      combat.prototypes[name],
		team:           team,
		level:          level,
		position:       position,
		items:          make([]ItemName, 0, 3),
		currentHealth:  int(math.Round(prototype.Health[level])),
		currentMana:    int(math.Round(prototype.StartMana)),
		report:         new(CombatReport),
		damagers:       make([]Unit, 0, 8),
		addModif:       make(map[ModifierName][]Modifier),
		multModif:      make(map[ModifierName][]Modifier),
		eventNotifiers: make(map[EventType][]func(*Event)),
		effects:        make(map[int]*Effect),
	}

	combat.Units = append(combat.Units, unit)
	combat.UnitRetriever[unit.id] = unit

	if unit.position >= 0 {
		combat.Board[unit.position] = unit
	}

	return unit
}

func CreateRandomGetNeighborsFunc() func(int) []int {
	foo := CreateGetNeighborsFunction(8, 7)
	return func(x int) []int {
		neighbors := foo(x)
		rand.Shuffle(len(neighbors), CreateIntSwapFunction(neighbors))
		return neighbors
	}
}

func incrementMapEntry(m map[string]int, e string) {
	_, ok := m[e]

	if ok {
		m[e]++
	} else {
		m[e] = 1
	}

}

func LogC(combat *Combat, message string) {
	combat.CombatLog = append(combat.CombatLog, message)
}

// TODO: Document that 0 results in instant combat
// TODO: Document that negative value results in  slower

func CombatSlowRun(combat *Combat, speed float64) {
	modifier := 0.0
	if speed != 0 {
		modifier = 1 / speed
		if speed < 0 {
			modifier = -speed
		}
	}

	combatProcessTick := StartCombat(combat)
	Display(combat)
	fmt.Println("Combat starting soon...")
	time.Sleep(time.Duration(time.Second))

	for {
		if combat.WinningTeam > 0 {
			return
		}

		start := time.Now()
		combatProcessTick(combat)
		Display(combat)
		elapsed := time.Since(start).Nanoseconds()

		fmt.Printf("%s\n\nt:%d, %fms\n%s", Grey, combat.CurrentTick, float64(elapsed)/1000000.0, Reset)

		// subtract time we have been computing
		// TODO: The order of this is wrong
		// It should be: compute, buffer display, wait, print

		if modifier != 0 {
			time.Sleep(time.Duration(1000000000*modifier/DefaultTickRate)*time.Nanosecond - time.Duration(elapsed))
		}

	}
}

func CreateUnitSwapFunction(arr []Unit) func(int, int) {
	return func(x int, y int) {
		arr[x], arr[y] = arr[y], arr[x]
	}
}
func CreateIntSwapFunction(arr []int) func(int, int) {
	return func(x int, y int) {
		arr[x], arr[y] = arr[y], arr[x]
	}
}

func (combat *Combat) DealEffect(options *EffectDescription) *Effect {
	effect := &Effect{
		Identity: combat.NextId(),
		Start:    combat.CurrentTick,

		Name:     options.name,
		Amount:   options.amount,
		Duration: options.duration,

		Target: options.target,
		Source: options.source,
	}
	effect.Target.AddEffect(effect)

	deleteEffect := func() {
		effect.Target.DeleteEffect(effect)
	}
	combat.scheduling.Schedule(deleteEffect, effect.Start+effect.Duration)
	return effect
}

func UnitMove(unit Unit, newPosition int) {
	// TODO: Make unit.Position private
	// Maybe the board too
	// TODO: Btw naming types is good, but actually not, because it could make code less readable

	if unit.Combat().Board[newPosition] != nil {
		panic("Trying to move to an empty position!")
	}

	unit.Combat().Board[unit.Position()] = nil
	unit.SetPosition(newPosition)

	if newPosition >= 0 {
		unit.Combat().Board[newPosition] = unit
	}
}

func DealHeal(source Unit, target Unit, heal int) {
	if !IsAlive(target) {
		return
	}

	// TODO: We can track healing lost through wound

	woundPercentage, author := GetActiveEffectSum(target, Wound)
	woundLoss := PercentageOf(heal, woundPercentage)
	heal -= woundLoss

	if author != nil {
		author.CombatReport().WoundDealt += woundLoss
	}
	overheal := 0

	if target.CurrentHealth()+heal > target.MaxHealth() {
		overheal = target.CurrentHealth() + heal - target.MaxHealth()
		target.Notify(NewEventPayload(OnOverheal, overheal))
	}

	LogC(source.Combat(), fmt.Sprintf("%s +[%d]+> %s", source.Name(), heal-overheal, target.Name()))

	target.AddHealth(heal - overheal)
	source.CombatReport().Healing += heal // TODO: Do we CombatReport() overheal in combat CombatReport()?
	target.Notify(NewEventPayload(OnHealthChange, heal))
}

func IsTeamEliminated(team TeamId, combat *Combat) bool {
	units := Filter(combat.Units, func(u Unit) bool {
		return u.Team() == team
	})
	for _, unit := range units {
		if unit.IsAlive() {
			return false
		}
	}
	return true
}

// Source eliminates target
func MaybeNotifyUnitElimination(combat *Combat, source Unit, target Unit) {
	if IsAlive(target) {
		return
	}
	// TODO: Log elimination here
	combat.CombatLog = append(combat.CombatLog, fmt.Sprintf("%s eliminates %s", source.Name(), target.Name()))
	combat.Board[target.Position()] = nil
	target.SetPosition(-1)

	for _, damagingUnit := range target.Damagers() {
		damagingUnit.Notify(NewEvent(OnTakedown))
	}

	// TODO: An optimization is to keep track of alive champions rather than
	// Manually check every single time

	if IsTeamEliminated(target.Team(), combat) {
		combat.WinningTeam = source.Team()
	}

	source.Notify(NewEvent(OnKill))
	target.Notify(NewEvent(OnDeath))
}

func UnitDamageResisted(unit Unit, damage int, resistance int) int {
	return Resistance(unit.ArmorResist(), damage) + PercentageOf(damage, unit.Durability())
}

func DealDamage(source Unit, target Unit, damageType DamageType, damage int) {

	isInvulnerable := GetActiveEffectPresenceMap(target)[Invulnerability]
	if isInvulnerable {
		// TODO: What do we do? Can this happen?
	}

	combat := source.Combat()
	if !target.IsAlive() {
		// TODO: log here is this line necessary
		return
	}

	damageResisted := 0

	if damageType == AttackDamageType {
		sunderPercentage, author := GetActiveEffectSum(target, Sunder)
		netResistance := PercentageOf(target.ArmorResist(), sunderPercentage)

		rawResisted := UnitDamageResisted(target, damage, target.ArmorResist())
		damageResisted = UnitDamageResisted(target, damage, netResistance)

		if author != nil {
			author.CombatReport().SunderDealt += rawResisted - damageResisted
		}

		source.CombatReport().AttackDamage += damage - damageResisted
		target.CombatReport().AttackResist += damageResisted
	} else if damageType == AbilityDamageType {
		shredPercentage, author := GetActiveEffectSum(target, Shred)
		netResistance := PercentageOf(target.ArmorResist(), shredPercentage)

		rawResisted := UnitDamageResisted(target, damage, target.MagicResist())
		damageResisted = UnitDamageResisted(target, damage, netResistance)

		if author != nil {
			author.CombatReport().ShredDealt += rawResisted - damageResisted
		}

		source.CombatReport().AttackDamage += damage - damageResisted
		target.CombatReport().AttackResist += damageResisted
	} else {

		// True Damage ignores resistance and durability
		source.CombatReport().TrueDamage += damage
	}

	shields := GetActiveEffects(target, Shield)
	sort.Slice(shields, func(x int, y int) bool {
		return EffectRemaining(shields[x], combat) < EffectRemaining(shields[y], combat)
	})

	damageBlocked := 0
	remainingDamage := MaxInt(damage-damageResisted, 0)

	for _, shield := range shields {
		currentShieldDamage := MinInt(remainingDamage, int(shield.Amount))

		shield.Amount -= float64(currentShieldDamage)

		if shield.Amount <= 0 {
			DeactivateEffectEarly(combat, shield)
		}

		remainingDamage -= currentShieldDamage
		damageBlocked += currentShieldDamage
	}

	if damageBlocked > 0 {
		source.Notify(NewEventPayload(OnDamageShield, damageBlocked))
	}

	target.AddHealth(-remainingDamage)

	target.CombatReport().ShieldsBlock += damageBlocked
	target.CombatReport().HealthBlock += remainingDamage

	target.AddDamager(source)

	source.Notify(NewEventPayload(OnDealingDamage, damage))
	target.Notify(NewEventPayload(OnReceivingDamage, damage))
	target.Notify(NewEventPayload(OnHealthChange, -damage))

	MaybeNotifyUnitElimination(combat, source, target)
}

func IsTargetInRange(unit Unit, combat *Combat, target Unit) bool {
	//if target.Position < 0 || unit.Position < 0 {
	//	return false TODO: The commented out code should not be necessary
	//}
	return combat.GetLogicalDistance(unit.Position(), target.Position()) <= unit.HexRange()
}

func TargetNewEnemy(combat *Combat, unit Unit) []int {
	unit.SetTarget(nil)
	unit.ResetChaseDistance()

	potentialTargets, previous := combat.TargetSelector(unit.Position(), combat, unit.Team())

	if len(potentialTargets) == 0 {
		// TODO: Log that there are no available paths to any enemy unit
		return make([]int, 0)
	}

	unit.SetTarget(potentialTargets[0])
	return MakePath(unit.Position(), unit.Target().Position(), previous)
}

func RunTargetingLogic(unit Unit) []int {
	combat := unit.Combat()

	path := make([]int, 0)
	if unit.Target() == nil || !IsAlive(unit.Target()) {
		path = TargetNewEnemy(combat, unit)

	} else if !IsTargetInRange(unit, combat, unit.Target()) {
		path = FindPath(unit.Position(), unit.Target().Position(), combat.Board, unit.Team(), combat.GetNeighbors)

		if len(path) == 0 {
			// Path to current target not found, change target
			// TODO: I feel like this could be simplified somehow
			path = TargetNewEnemy(combat, unit)
		} else {

			chaseHex := path[MinInt(unit.ChaseDistance(), len(path)-1)]
			postChaseRangeDeficit := combat.GetLogicalDistance(chaseHex, unit.Target().Position()) - unit.HexRange()

			if postChaseRangeDeficit > 0 || unit.ChaseDistance() >= MaximumChaseDistance {
				path = TargetNewEnemy(combat, unit)
			}
		}
	}
	return path
}

func IsCasting(unit Unit) bool {
	return unit.StateName() == CastAnimationState
}

func NextAttackDamage(unit Unit) int {
	isCritical := rand.Float64() < unit.CriticalStrikeChance()
	critMod := 1.0
	if isCritical {
		critMod = unit.CriticalStrikeDamage()
	}
	attackDamageF := float64(unit.AttackDamage()) * critMod * (1 + unit.DamageAmp())
	return RoundUpOptimistically(attackDamageF)
}

func UnitProcessTick(unit Unit, combat *Combat) bool {
	// TODO: Do adaptive helm and start combat mana testing on melee abilities

	// Phase 0: Return if we do not process this Unit
	// 			Ex: !IsAlive, Stun, Held,

	// Phase 1: If Unit can cast, we cast

	// Phase 2: Determine if Unit has an current or potential Target in range
	//			If not, interrupt the interruptibles and delegate this Champion to path logic

	// Phase 3:
	//
	return false
}

func CooperativePathingTick(unit []Unit, combat *Combat) {

}

func RegisterItems(champion Unit) {
	for _, item := range champion.Items() {
		RegisterItem(item, false, champion)
	}
}

func StartCombat(combat *Combat) func(*Combat) {
	for _, unit := range combat.Units {
		// TODO: This line should really not be necessary
		combat.UnitRetriever[unit.Id()] = unit
	}

	for _, unit := range combat.Units {

		// RegisterAbility()
		// RegisterItems() TODO: Register items here
		// TODO: RegisterTraitsUnit
		// TODO: RegisterAugmentsUnit

		// TODO: RegisterAnomaly

		// TODO: Document following line. Maybe a better way to do this?

		unit.SetHealth(unit.MaxHealth())

		// TODO: Register Traits here based on the intersection of Unit's traits and Team's active traits
		// TODO: RegisterTrait(unit, combat)
		unit.Notify(NewEvent(OnCombatStart))
	}
	return func(combat *Combat) { // ProcessTick(combat)

		if combat.WinningTeam > 0 {
			// TODO: do gameover somehow differently
			return
		}

		// Execute all scheduled calls!
		combat.scheduling.Dispatch()

		rand.Shuffle(len(combat.Units), CreateUnitSwapFunction(combat.Units))

		for _, unit := range combat.Units {
			// TODO: Only notify if the tick is correct
			unit.Notify(NewEventPayload(OnTick, combat.CurrentTick))
		}

		melee := Filter(combat.Units, func(u Unit) bool { return u.HexRange() <= 1 })
		ranged := Filter(combat.Units, func(u Unit) bool { return u.HexRange() > 1 })

		for _, champion := range melee {

			UnitProcessTick(champion, combat)
		}

		for _, champion := range ranged {
			UnitProcessTick(champion, combat)
		}

		combat.CurrentTick++
	}
}
