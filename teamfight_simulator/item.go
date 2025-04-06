package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

// TODO: One unit test could be one giga unit with all items except for TG
// Second unit test could be units with 3 copies of item, if item is not unique and not base

func IsBaseItem(item ItemName) bool {
	switch item {
	case BFSword:
		return true
	case ChainVest:
		return true
	case GiantsBelt:
		return true
	case NeedlesslyLargeRod:
		return true
	case NegatronCloak:
		return true
	case RecurveBow:
		return true
	case SparringGloves:
		return true
	case TearOfTheGoddess:
		return true
	default:
		return false
	}
}

func IsUniqueItem(item ItemName) bool {
	switch item {
	case BlueBuff:
		return true
	case EdgeOfNight:
		return true
	case LastWhisper:
		return true
	case Morellonomicon:
		return true
	case ThiefsGloves:
		return true
	default:
		return false
	}
}

func GetBaseItems() []ItemName {
	return []ItemName{
		BFSword,
		ChainVest,
		GiantsBelt,
		NeedlesslyLargeRod,
		NegatronCloak,
		RecurveBow,
		SparringGloves,
		TearOfTheGoddess,
	}
}

func GetCraftableItems() []ItemName {
	return []ItemName{
		AdaptiveHelm,
		ArchangelsStaff,
		Bloodthirster,
		BlueBuff,
		BrambleVest,
		Crownguard,
		Deathblade,
		DragonsClaw,
		EdgeOfNight,
		Evenshroud,
		GargoyleStoneplate,
		GiantSlayer,
		Guardbreaker,
		GuinsoosRageblade,
		HandsOfJustice,
		HextechGunblade,
		InfinityEdge,
		IonicSpark,
		JeweledGauntlet,
		LastWhisper,
		Morellonomicon,
		NashorsTooth,
		ProtectorsVow,
		Quicksilver,
		RabadonsDeathcap,
		RedBuff,
		Redemption,
		RunaansHurricane,
		SpearOfShojin,
		StatikkShiv,
		SteadfastHeart,
		SteraksGage,
		SunfireCape,
		ThiefsGloves,
		TitansResolve,
		WarmogsArmor,
	}
}

type Item interface {
	// TODO: Redo items into this interface + struct
	Register()
	Data() map[string]float64
}

func CraftItem(item1 ItemName, item2 ItemName) ItemName {
	craft := func(i1 ItemName, i2 ItemName) string {
		if i1 > i2 {
			i1, i2 = i2, i1
		}
		return ItemToString(i1) + "|" + ItemToString(i2)
	}

	switch craft(item1, item2) {
	case craft(BFSword, BFSword):
		return Deathblade
	case craft(BFSword, ChainVest):
		return EdgeOfNight
	case craft(BFSword, GiantsBelt):
		return SteraksGage
	case craft(BFSword, NeedlesslyLargeRod):
		return HextechGunblade
	case craft(BFSword, NegatronCloak):
		return Bloodthirster
	case craft(BFSword, RecurveBow):
		return GiantSlayer
	case craft(BFSword, SparringGloves):
		return InfinityEdge
	case craft(BFSword, TearOfTheGoddess):
		return SpearOfShojin

	case craft(ChainVest, ChainVest):
		return BrambleVest
	case craft(ChainVest, GiantsBelt):
		return SunfireCape
	case craft(ChainVest, NeedlesslyLargeRod):
		return Crownguard
	case craft(ChainVest, NegatronCloak):
		return GargoyleStoneplate
	case craft(ChainVest, RecurveBow):
		return TitansResolve
	case craft(ChainVest, SparringGloves):
		return Guardbreaker
	case craft(ChainVest, TearOfTheGoddess):
		return SteadfastHeart

	case craft(GiantsBelt, GiantsBelt):
		return WarmogsArmor
	case craft(GiantsBelt, NeedlesslyLargeRod):
		return Morellonomicon
	case craft(GiantsBelt, NegatronCloak):
		return AdaptiveHelm
	case craft(GiantsBelt, RecurveBow):
		return RedBuff
	case craft(GiantsBelt, SparringGloves):
		return ProtectorsVow
	case craft(GiantsBelt, TearOfTheGoddess):
		return Redemption

	case craft(NeedlesslyLargeRod, NeedlesslyLargeRod):
		return RabadonsDeathcap
	case craft(NeedlesslyLargeRod, NegatronCloak):
		return IonicSpark
	case craft(NeedlesslyLargeRod, RecurveBow):
		return GuinsoosRageblade
	case craft(NeedlesslyLargeRod, SparringGloves):
		return JeweledGauntlet
	case craft(NeedlesslyLargeRod, TearOfTheGoddess):
		return ArchangelsStaff

	case craft(NegatronCloak, NegatronCloak):
		return DragonsClaw
	case craft(NegatronCloak, RecurveBow):
		return RunaansHurricane
	case craft(NegatronCloak, SparringGloves):
		return Quicksilver
	case craft(NegatronCloak, TearOfTheGoddess):
		return Evenshroud

	case craft(RecurveBow, RecurveBow):
		return NashorsTooth
	case craft(RecurveBow, SparringGloves):
		return LastWhisper
	case craft(RecurveBow, TearOfTheGoddess):
		return StatikkShiv

	case craft(SparringGloves, SparringGloves):
		return ThiefsGloves
	case craft(SparringGloves, TearOfTheGoddess):
		return HandsOfJustice

	case craft(TearOfTheGoddess, TearOfTheGoddess):
		return BlueBuff

	default:
		panic("Unimplemented item or o1 chatgpt mistake")
	}
}

func ContainsItem(unit Unit, item ItemName) bool {
	for _, item2 := range unit.Items() {
		if item2 == item {
			return true
		}
	}
	return false
}

// Once you add an item to a unit
// You need to do registrations over again
func TryAddUnitItem(unit Unit, item ItemName) bool {
	if len(unit.Items()) == 3 {
		return false
	}

	if IsUniqueItem(item) && ContainsItem(unit, item) {
		return false
	}

	if len(unit.Items()) == 0 {
		unit.AddItem(item)
		return true
	} else if item == ThiefsGloves {
		// Thief's Gloves must be the only item
		return false
	}

	lastIndex := len(unit.Items()) - 1
	lastItem := unit.Items()[lastIndex]

	if IsBaseItem(item) && IsBaseItem(lastItem) {

		if CraftItem(item, lastItem) == ThiefsGloves && len(unit.Items()) > 1 {
			// Thief's gloves must be the only item
			return false
		}

		unit.Items()[lastIndex] = CraftItem(lastItem, item)
		return true
	}

	unit.AddItem(item)
	return true
}

/*****      Base Items      *****/
func RegisterBFSword(unit Unit) {
	baseAttackDamage := 10.0
	unit.AddModifier(AdditiveModifier(AttackDamage).WithConstant(baseAttackDamage))

}
func RegisterChainVest(unit Unit) {
	baseArmorResist := 20.0
	unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(baseArmorResist))
}
func RegisterNeedlesslyLargeRod(unit Unit) {
	baseAbilityPower := 10.0
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(baseAbilityPower))
}
func RegisterNegatronCloak(unit Unit) {
	baseMagicResist := 20.0
	unit.AddModifier(AdditiveModifier(MagicResist).WithConstant(baseMagicResist))
}
func RegisterGiantsBelt(unit Unit) {
	baseHealth := 150.0
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(baseHealth))
}
func RegisterRecurveBow(unit Unit) {
	baseAttackspeed := 0.10
	unit.AddModifier(MultiplicativeModifier(AttackSpeed).WithConstant(baseAttackspeed))
}
func RegisterSparringGloves(unit Unit) {
	baseCriticalChance := 0.20
	unit.AddModifier(AdditiveModifier(CriticalStrikeChance).WithConstant(baseCriticalChance))
}
func RegisterTearOfTheGoddess(unit Unit) {
	baseStartMana := 15.0
	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(baseStartMana))
}

/*****      Craftable Items      *****/
func RegisterAdaptiveHelm(unit Unit, isRadiant bool) {
	// TODO: Have an object+functions for cooldowns??

	// TODO: Ensure that these three get added regardless of backline/frontline

	d := map[string]float64{
		"start_mana":    15.0,
		"magic_resist":  20.0,
		"ability_power": 15.0,
	}

	manaGainOnHit := 1
	manaGainOnInterval := 10

	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(d["start_mana"]))
	unit.AddModifier(AdditiveModifier(MagicResist).WithConstant(d["magic_resist"]))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(d["ability_power"]))

	frontlineArmorResist := 40.0
	frontlineMagicResist := 40.0

	backlineManaMaxCooldown := SecondsToTicks(3)
	cooldown := backlineManaMaxCooldown // TODO: Set to 0 if adaptive helm instantly procs
	backlineAbilityPower := 20.0

	unit.RegisterEvent(OnCombatStart, func(e *Event) {
		if unit.IsFrontStarter() {
			unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(frontlineArmorResist))
			unit.AddModifier(AdditiveModifier(MagicResist).WithConstant(frontlineMagicResist))
			unit.RegisterEvent(OnAttacked, func(e *Event) {
				unit.AddMana(manaGainOnHit) // TODO: Should we have onManaChange and a method here
			})
		} else if unit.IsBackStarter() {
			unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(backlineAbilityPower))
			unit.RegisterEvent(OnTick, func(e *Event) {
				cooldown--
				if TickRoundingScheme(backlineManaMaxCooldown) <= 0 {
					cooldown = backlineManaMaxCooldown
					unit.AddMana(manaGainOnInterval)
				}
			})
		} else {
			panic("Unit start row is not recognized!")
		}
	})
}

// TODO: all AddModifier AdditiveModifier()and AddModifier MultiplicativeModifier()has to be changed
// Also figure out what to do with Data / Params observability??????
// Data / Params maybe shouldnt be with modifier, it maybe should be with Item
func RegisterArchangelsStaff(unit Unit, isRadiant bool) {
	startMana := 15.0
	abilityPower := 20.0

	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(startMana))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(abilityPower))

	abilityPowerIncrement := 30
	interval := SecondsToTicks(5)

	// Description says on Combat Start, but I'll just make it always active
	unit.AddModifier(AdditiveModifier(AbilityPower).WithRetriever(
		func(Unit) float64 {
			stacks := RoundDown(float64(unit.Combat().CurrentTick) / interval)
			return float64(stacks * abilityPowerIncrement)
		}))
}
func RegisterBloodthirster(unit Unit, isRadiant bool) {
	attackDamage := .20
	magicResist := 20.0
	abilityPower := 20.0
	omnivamp := 0.20

	activatePercentage := 0.4
	shieldPercentage := 0.25
	shieldDuration := SecondsToTicksRound(5)

	unit.AddModifier(AdditiveModifier(AttackDamage).WithConstant(attackDamage))
	unit.AddModifier(AdditiveModifier(MagicResist).WithConstant(magicResist))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(abilityPower))
	unit.AddModifier(AdditiveModifier(Omnivamp).WithConstant(omnivamp))

	isActivated := false
	unit.RegisterEvent(OnHealthChange, func(e *Event) {
		if isActivated || unit.CurrentHealth() > HealthPercentage(unit, activatePercentage) {
			return
		}
		isActivated = true

		shieldAmount := float64(HealthPercentage(unit, shieldPercentage))
		unit.DealEffect(
			EffectBuilder(Shield).
				Source(unit).
				Target(unit).
				Duration(shieldDuration).
				Amount(shieldAmount))
	})
}

func RegisterBlueBuff(unit Unit, isRadiant bool) {
	combat := unit.Combat()

	startMana := 20.0
	abilityPower := 20.0
	attackDamage := .20

	activeDamageAmp := 0.08
	activeDuration := SecondsToTicksRound(8)

	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(startMana))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(abilityPower))
	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(attackDamage))

	unit.RegisterEvent(OnTakedown, func(e *Event) {
		expirationTick := combat.CurrentTick + activeDuration
		unit.AddModifier(AdditiveModifier(DamageAmp).WithConstant(activeDamageAmp).WithExpiration(expirationTick))
	})
}
func RegisterBrambleVest(unit Unit, isRadiant bool) {
	armorResist := 55.0
	maxHealth := 0.05
	durability := 0.08

	unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(armorResist))
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(maxHealth))
	unit.AddModifier(AdditiveModifier(Durabiltiy).WithConstant(durability))
	brambleRange := 1
	cooldownDuration := SecondsToTicksRound(2)
	lastActivated := -cooldownDuration

	// After any attack to all adjaceny enemies
	// To all adjacent with 2 second cooldown
	magicDamage := 100

	unit.RegisterEvent(OnReceivingDamage, func(e *Event) {
		if unit.Combat().CurrentTick-lastActivated < cooldownDuration {
			return
		}
		LogC(unit.Combat(), "Bramble vest proc by "+unit.Name())
		lastActivated = unit.Combat().CurrentTick
		adjacent := unit.SearchUnits(Enemies(unit).WithDistanceLimit(brambleRange))
		for _, target := range adjacent {
			if target.Team() != unit.Team() {
				DealDamage(unit, target, AbilityDamageType, magicDamage)
			}
		}
	})
}
func RegisterCrownguard(unit Unit, isRadiant bool) {
	combat := unit.Combat()

	ap := 20.0
	ar := 20.0
	hp := 100.0

	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(ap))
	unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(ar))
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(hp))

	shieldPercentage := 0.30
	shieldDuration := SecondsToTicksRound(8)
	abilityPowerOnExpiration := 35.0

	unit.RegisterEvent(OnCombatStart, func(e *Event) {
		shieldAmount := float64(HealthPercentage(unit, shieldPercentage))
		unit.DealEffect(
			EffectBuilder(Shield).
				Source(unit).
				Target(unit).
				Duration(shieldDuration).
				Amount(shieldAmount))
	})

	unit.AddModifier(AdditiveModifier(AbilityPower).WithRetriever(
		func(u Unit) float64 {
			if combat.CurrentTick < shieldDuration {
				return 0.0
			}
			return abilityPowerOnExpiration
		}))
}

func RegisterDeathblade(unit Unit, isRadiant bool) {
	attackDamage := 0.66
	bonusDamage := .07 // What the heck is this??? TODO: Check if this is damage amp
	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(attackDamage))
	unit.AddModifier(AdditiveModifier(DamageAmp).WithConstant(bonusDamage))
}

func RegisterDragonsClaw(unit Unit, isRadiant bool) {
	maxHealth := 0.09
	healPercentage := 0.05

	cooldown := 0.0
	maxCooldown := SecondsToTicks(2)

	unit.AddModifier(MultiplicativeModifier(MaximumHealth).WithConstant(maxHealth))
	unit.RegisterEvent(OnTick, func(e *Event) {
		if !IsAlive(unit) {
			// TODO: maybe have a deregister function that can be called??
			return
		}
		if cooldown > 0 {
			cooldown -= 1
		}
		cooldown = maxCooldown
		DealHeal(unit, unit, HealthPercentage(unit, healPercentage))
	})
}

func RegisterEdgeOfNight(unit Unit, isRadiant bool) {
	attackDamage := 0.10
	armorResist := 20.0

	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(attackDamage))
	unit.AddModifier(AdditiveModifier(AttackDamage).WithConstant(armorResist))

	// TODO: We come back to this later because we need to
	// Implement untargetability and shedding negative effects
	// Also we need to figure out the attack speed thing

	panic("unimplemented")
}

func ApplyHexRangeDebuff(unit Unit, affectedEnemies map[int]*Effect, effectType EffectName, effectDistance int, effectAmount float64, effectTextId string) {
	combat := unit.Combat()
	longDuration := SecondsToTicksRound(3600)

	unit.RegisterEvent(OnTick, func(e *Event) {
		toSunder := unit.SearchUnits(
			Enemies(unit).WithDistanceLimit(effectDistance).WithByRandom(),
		)

		for _, enemy := range toSunder {
			_, exists := affectedEnemies[enemy.Id()]
			if exists {
				continue
			}
			effect := unit.DealEffect(
				EffectBuilder(Shield).
					Source(unit).
					Target(unit).
					Duration(longDuration).
					Amount(effectAmount))
			affectedEnemies[enemy.Id()] = effect
		}

		for enemyId, effect := range affectedEnemies {

			enemy := combat.UnitRetriever[enemyId]

			if !IsAlive(enemy) {
				continue
			}
			if combat.GetLogicalDistance(unit.Position(), enemy.Position()) > effectDistance {
				DeactivateEffectEarly(combat, effect)
				delete(affectedEnemies, enemyId)
			}
		}
	})

	unit.RegisterEvent(OnDeath, func(e *Event) {
		for _, effect := range affectedEnemies {
			DeactivateEffectEarly(combat, effect)
		}

		// Just in case this unit is revived somehow!
		affectedEnemies = make(map[int]*Effect)
	})
}

func RegisterEvenshroud(unit Unit, isRadiant bool) {

	sunderAmount := 0.30
	sunderDistance := 2

	armorResist := 25.0
	magicResist := 25.0
	resistDuration := SecondsToTicksRound(10)

	effectTextId := "Evenshroud"

	removeArmorResist := unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(armorResist))
	removeMagicResist := unit.AddModifier(AdditiveModifier(MagicResist).WithConstant(magicResist))

	expirationTick := unit.Combat().CurrentTick + resistDuration
	unit.Combat().scheduling.Schedule(removeArmorResist, expirationTick)
	unit.Combat().scheduling.Schedule(removeMagicResist, expirationTick)

	sunderedEnemies := make(map[int]*Effect)
	ApplyHexRangeDebuff(unit, sunderedEnemies, Sunder, sunderDistance, sunderAmount, effectTextId)
}

func RegisterGargoyleStoneplate(unit Unit, isRadiant bool) {
	magicResist := 30.0
	armorResist := 30.0
	hp := 100.0

	mr_increment := 10.0
	ar_increment := 10.0

	unit.AddModifier(AdditiveModifier(ArmorResist).WithRetriever(
		func(u Unit) float64 {
			bonus := 0.0
			for _, otherUnit := range unit.Combat().Units {
				if otherUnit.Target().Id() == unit.Id() {
					bonus += ar_increment
				}
			}
			return armorResist + bonus
		}))

	unit.AddModifier(AdditiveModifier(MagicResist).WithRetriever(func(u Unit) float64 {
		bonus := 0.0
		for _, otherUnit := range unit.Combat().Units {
			if otherUnit.Target().Id() == unit.Id() {
				bonus += mr_increment
			}
		}
		return magicResist + bonus
	}))

	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(hp))

}
func RegisterGiantSlayer(unit Unit, isRadiant bool) {
	enemyHealthRequirement := 1750
	ad := .30
	as := .10
	ap := 10.0
	bonusDamage := 0.25

	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(ad))
	unit.AddModifier(MultiplicativeModifier(AttackSpeed).WithConstant(as))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(ap))

	unit.AddModifier(AdditiveModifier(DamageAmp).WithRetriever(
		func(u Unit) float64 {
			// TODO: This may not quite be right
			// TODO: What happen's with Ruunan's hurricane + this??
			if unit.Target().MaxHealth() > enemyHealthRequirement {
				// TODO: Does the 25% damage bonus also multiply the damage amp bonus damage?
				return bonusDamage
			}
			return 0
		}))
}

func RegisterGuardbreaker(unit Unit, isRadiant bool) {
	hp := 150.0
	critChance := 0.20
	ap := 10.0
	as := 0.20

	postShieldAmp := 0.25
	ampDuration := SecondsToTicksRound(3)
	lastDamageShields := -ampDuration - 1

	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(hp))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(ap))
	unit.AddModifier(MultiplicativeModifier(AttackSpeed).WithConstant(as))
	unit.AddModifier(AdditiveModifier(CriticalStrikeChance).WithConstant(critChance))

	// TODO: Does the 25% damage bonus also multiply the damage amp bonus damage?
	unit.AddModifier(AdditiveModifier(DamageAmp).WithRetriever(
		func(u Unit) float64 {
			if unit.Combat().CurrentTick-lastDamageShields >= ampDuration {
				return 0.0
			}
			return postShieldAmp
		}))

	unit.RegisterEvent(OnDamageShield, func(e *Event) {
		lastDamageShields = unit.Combat().CurrentTick
	})

}
func RegisterGuinsoosRageblade(unit Unit, isRadiant bool) {

	stacks := 0.0

	baseAttackSpeed := 0.15
	stackAttackSpeed := 0.05
	ap := 10.0

	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(ap))

	unit.AddModifier(MultiplicativeModifier(AttackSpeed).WithRetriever(func(u Unit) float64 {
		return baseAttackSpeed + stackAttackSpeed*stacks
	}))

	unit.RegisterEvent(OnAttack, func(e *Event) {
		stacks++
	})
}

func RegisterHandsOfJustice(unit Unit, isRadiant bool) {

	startMana := 15.0
	critChance := 0.20

	// Effect 0
	ad := 0.15
	ap := 15.0

	// Effect 1
	omnivamp := 0.15

	effectMultipliers := []float64{1.0, 1.0}

	unit.RegisterEvent(OnCombatStart, func(e *Event) {
		doublingIndex := rand.IntN(2)
		otherIndex := (doublingIndex + 1) % len(effectMultipliers)
		effectMultipliers[doublingIndex] = 2.0
		effectMultipliers[otherIndex] = 1.0
	})

	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(startMana))
	unit.AddModifier(AdditiveModifier(CriticalStrikeChance).WithConstant(critChance))

	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithRetriever(
		func(u Unit) float64 {
			return ad * effectMultipliers[0]
		}))

	unit.AddModifier(AdditiveModifier(AbilityPower).WithRetriever(
		func(u Unit) float64 {
			return ap * effectMultipliers[0]
		},
	))

	unit.AddModifier(AdditiveModifier(Omnivamp).WithRetriever(
		func(u Unit) float64 {
			return omnivamp * effectMultipliers[1]
		},
	))
}

func RegisterHextechGunblade(unit Unit, isRadiant bool) {
	combat := unit.Combat()
	ad := 0.15
	ap := 15.0
	omnivamp := 0.20
	healingPercentage := 0.20

	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(ad))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(ap))
	unit.AddModifier(AdditiveModifier(Omnivamp).WithConstant(omnivamp))

	unit.RegisterEvent(OnDealingDamage, func(e *Event) {
		damage := float64(e.Payload[0])
		var lowestHealthAlly Unit = nil
		lowestHealthPercentage := 1.0

		for _, otherUnit := range combat.Units {
			if otherUnit.Team() != unit.Team() || otherUnit.Id() == unit.Id() || !IsAlive(otherUnit) {
				continue
			}

			hpNow := float64(otherUnit.CurrentHealth())
			hpMax := float64(otherUnit.MaxHealth())
			otherHealthPercentage := hpNow / hpMax
			if otherHealthPercentage < lowestHealthPercentage {
				lowestHealthPercentage = otherHealthPercentage
				lowestHealthAlly = otherUnit
			}
		}

		if lowestHealthAlly != nil {
			healingAmount := RoundUpOptimistically(damage * healingPercentage)
			DealHeal(unit, lowestHealthAlly, healingAmount)
		}
	})

}

func RegisterInfinityEdge(unit Unit, isRadiant bool) {
	ad := 0.35
	critChance := 0.35
	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(ad))
	unit.AddModifier(AdditiveModifier(CriticalStrikeChance).WithRetriever(func(unit Unit) float64 {
		bonus := 0.0

		// TODO: Ambushers can critically strike
		// TODO: Does bad luck protection
		// TODO: IE and JG can critically strike
		// TODO: Piercing Lotus I
		// TODO: Piercing Lotus II
		// TODO: If abilities already critically strike without this item, set bonus crit chance to 10% (More if radiant?)

		return critChance + bonus
	}))

}

func RegisterIonicSpark(unit Unit, isRadiant bool) {
	ap := 15.0
	mr := 25.0
	hp := 150.0

	manaDamageMultiplier := 1.6

	shredDistance := 2
	shredAmount := 0.30

	textEffectId := "IonicSpark"

	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(ap))
	unit.AddModifier(AdditiveModifier(MagicResist).WithConstant(mr))
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(hp))

	unit.RegisterEvent(OnEnemyAbilityCast, func(event *Event) {
		enemyId := event.Payload[0]
		enemy := unit.Combat().UnitRetriever[enemyId]
		enemyMaxMana := float64(enemy.MaxMana())
		magicDamage := RoundUpOptimistically(enemyMaxMana * manaDamageMultiplier)
		DealDamage(unit, enemy, AbilityDamageType, magicDamage)
	})

	shreddedEnemies := make(map[int]*Effect)
	ApplyHexRangeDebuff(unit, shreddedEnemies, Shred, shredDistance, shredAmount, textEffectId)
}

func RegisterJeweledGauntlet(unit Unit, isRadiant bool) {
	ap := 35.0
	critChance := 0.35
	unit.AddModifier(MultiplicativeModifier(AbilityPower).WithConstant(ap))
	unit.AddModifier(AdditiveModifier(CriticalStrikeChance).WithRetriever(
		func(unit Unit) float64 {
			bonus := 0.0

			// TODO: Ambushers can critically strike
			// TODO: Does bad luck protection
			// TODO: IE and JG can critically strike
			// TODO: Piercing Lotus I
			// TODO: Piercing Lotus II
			// TODO: If abilities already critically strike without this item, set bonus crit chance to 10% (More if radiant?)

			return critChance + bonus
		}))
}

func RegisterLastWhisper(unit Unit, isRadiant bool) {
	// TODO: Have to apply Morello's on DealDamage, but only if the damage is physical
	// (And maybe more places if an ability doesn't do damage???)

	panic("unimplemented")
}
func RegisterMorellonomicon(unit Unit, isRadiant bool) {
	// TODO: Have to apply Morello's on DealDamage.
	// (And maybe more places if an ability doesn't do damage???)

	panic("unimplemented")
}
func RegisterNashorsTooth(unit Unit, isRadiant bool) {
	combat := unit.Combat()

	nashorsAttackSpeed := 0.40
	nashorsDuration := 5

	unit.RegisterEvent(OnCastEnd, func(e *Event) {
		removeModifier := unit.AddModifier(MultiplicativeModifier(AttackSpeed).WithConstant(nashorsAttackSpeed))
		expirationTick := combat.CurrentTick + SecondsToTicksRound(nashorsDuration)
		combat.scheduling.Schedule(removeModifier, expirationTick)
	})

	panic("unimplemented")
}

func RegisterProtectorsVow(unit Unit, isRadiant bool) {
	ar := 20.0
	startMana := 30.0
	unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(ar))
	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(startMana))

	activationPercentage := 0.4
	shieldPercentage := 0.25
	shieldDuration := SecondsToTicksRound(5)
	arGain := 20.0
	mrGain := 20.0
	isActivated := false

	unit.RegisterEvent(OnHealthChange, func(e *Event) {
		activationHealth := HealthPercentage(unit, activationPercentage)
		shieldAmount := float64(HealthPercentage(unit, shieldPercentage))

		unit.DealEffect(
			EffectBuilder(Shield).
				Source(unit).
				Target(unit).
				Duration(shieldDuration).
				Amount(shieldAmount))

		if isActivated || unit.CurrentHealth() > activationHealth || !IsAlive(unit) {
			return
		}
		isActivated = true

		unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(arGain))
		unit.AddModifier(AdditiveModifier(MagicResist).WithConstant(mrGain))

	})
}
func RegisterQuicksilver(unit Unit, isRadiant bool) {
	// TODO: Have to implement immunity to crowd control, as an effect???
	panic("unimplemented")
}
func RegisterRabadonsDeathcap(unit Unit, isRadiant bool) {
	ap := 50.0
	bonusDamage := 0.2
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(ap))
	unit.AddModifier(AdditiveModifier(DamageAmp).WithConstant(bonusDamage))
}
func RegisterRedBuff(unit Unit, isRadiant bool) {
	panic("unimplemented")
}

func RegisterRedemption(unit Unit, isRadiant bool) {
	hp := 150.0
	startMana := 15.0
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(hp))
	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(startMana))

	healMissingPercentage := 0.15
	durabilityAmount := 0.10
	durabilityMaxCooldown := SecondsToTicksRound(5)
	durabilityDuration := SecondsToTicksRound(5)
	redemptionRange := 1
	cooldown := durabilityMaxCooldown

	// TODO: Does it proc on combat start?
	unit.RegisterEvent(OnTick, func(e *Event) {
		cooldown -= 1
		if cooldown > 0 {
			return
		}

		adjacentUnits := unit.SearchUnits(
			Friends(unit).WithDistanceLimit(redemptionRange),
		)

		for _, adjacentUnit := range adjacentUnits {
			if adjacentUnit.Team() == unit.Team() {

				missingHealth := float64(adjacentUnit.MissingHealth())
				healAmount := RoundUpOptimistically(missingHealth * healMissingPercentage)

				DealHeal(unit, adjacentUnit, healAmount)

				expirationTick := unit.Combat().CurrentTick + durabilityDuration
				// Unique: We're adding a modifier to a different unit
				removeDurability := adjacentUnit.AddModifier(AdditiveModifier(Durabiltiy).WithConstant(durabilityAmount))
				unit.Combat().scheduling.Schedule(removeDurability, expirationTick)
			}
		}
	})
}
func RegisterRunaansHurricane(unit Unit, isRadiant bool) {
	// TODO: Will come back later here. Complexities are:
	// 1. Bolts do not apply on-hit effects (??? Investigate) (Ex: Bramble Vest)
	// 2. Ranged Champions gain mana instantly with Runaans
	//    Otherwise, Ranged gains mana when projectile hits target
	//    But we do not even have a projectile system yet!!!
	// 3. Does Damage Amp apply to Runaan's bolts? It says 55% of attack damage
	// 4. Does this damage critically strike?
	panic("unimplemented")
}

func RegisterSpearOfShojin(unit Unit, isRadiant bool) {
	attackDamage := .20
	startMana := 15.0
	abilityPower := 20.0
	manaOnAttack := 5

	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(attackDamage))
	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(startMana))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(abilityPower))

	unit.RegisterEvent(OnAttack, func(e *Event) {
		unit.AddMana(manaOnAttack)
	})
}

func RegisterStatikkShiv(unit Unit, isRadiant bool) {
	// starting from the current target, it bounces 3 times to the enemy farthest from the previous one
	// TODO: Investigate if that's the correct pattern

	baseAttackSpeed := .20
	baseStartMana := 15.0
	baseAbilityPower := 15.0

	unit.AddModifier(MultiplicativeModifier(AttackSpeed).WithConstant(baseAttackSpeed))
	unit.AddModifier(AdditiveModifier(StartMana).WithConstant(baseStartMana))
	unit.AddModifier(AdditiveModifier(AbilityPower).WithConstant(baseAbilityPower))

	shredCount := 4
	shredDuration := SecondsToTicksRound(5)
	shredAmount := 0.30
	magicDamageAmount := 35
	frequency := 3

	attackCount := 0

	unit.RegisterEvent(OnAttack, func(event *Event) {
		combat := unit.Combat()
		attackCount++
		if attackCount%frequency != 0 {
			return
		}

		targetId := event.Payload[0]
		target := combat.UnitRetriever[targetId]

		alreadyShredded := make([]int, 0, shredCount)

		for i := 0; i < shredCount; i++ {

			// TODO: Shred and then deal damage of deal damage and then shred?

			unit.DealEffect(
				EffectBuilder(Shred).
					Source(unit).
					Target(target).
					Duration(shredDuration).
					Amount(shredAmount))

			DealDamage(unit, target, AbilityDamageType, magicDamageAmount)
			alreadyShredded = append(alreadyShredded, targetId)

			enemies := unit.SearchUnits(Enemies(unit))

			if len(enemies) == 0 || SliceContains(alreadyShredded, LastUnit(enemies).Id()) {
				// if there are no enemies to shred or if we already shredded the next target
				// we are done applying the statikk shiv
				break
			}
			target = LastUnit(enemies)
		}
	})
}

func RegisterSteadfastHeart(unit Unit, isRadiant bool) {
	baseArmorResist := 20.0
	baseCritChance := .20
	baseHealth := 200.0
	activationPercentage := 0.50

	inactiveDur := 0.08
	activeDur := 0.15

	unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(baseArmorResist))
	unit.AddModifier(AdditiveModifier(CriticalStrikeChance).WithConstant(baseCritChance))
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(baseHealth))

	unit.AddModifier(AdditiveModifier(Durabiltiy).WithRetriever(func(u Unit) float64 {

		if unit.CurrentHealth() > HealthPercentage(unit, activationPercentage) {
			return inactiveDur
		} else {
			return activeDur
		}
	}))

}

func RegisterSteraksGage(unit Unit, isRadiant bool) {
	baseHealth := 200.0
	baseAttackDamage := .15
	activationPercentage := 0.60
	recoveryPercentage := 0.25
	activedAttackDamageMultiplier := .35

	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(baseHealth))
	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(baseAttackDamage))

	isActivated := false

	unit.RegisterEvent(OnHealthChange, func(e *Event) {
		activationHealth := HealthPercentage(unit, activationPercentage)
		recoveryHealth := HealthPercentage(unit, recoveryPercentage)

		if isActivated || unit.CurrentHealth() > activationHealth || !IsAlive(unit) {
			return
		}

		isActivated = true

		unit.AddHealth(recoveryHealth)
		unit.AddModifier(MultiplicativeModifier(AttackDamage).WithConstant(activedAttackDamageMultiplier))
	})
}

func RegisterSunfireCape(unit Unit, isRadiant bool) {
	baseHealth := 250.0
	baseArmorResist := 20.0

	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(baseHealth))
	unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(baseArmorResist))

	// TODO: Figure out how the Overtime mode works
	// TODO: In regards to items that activate every X seconds

	cooldown := 0.0
	maximumCooldown := SecondsToTicks(2)

	capeDuration := SecondsToTicksRound(10)
	burnAmount := 0.01
	woundAmount := 0.33
	sunfireCapeRange := 2

	unit.RegisterEvent(OnTick, func(event *Event) {
		combat := unit.Combat()
		if TickRoundingScheme(cooldown) > 0 {
			cooldown--
			return
		}

		applicableEnemies := unit.SearchUnits(
			Enemies(unit).WithDistanceLimit(sunfireCapeRange).WithByRandom(),
		)

		if len(applicableEnemies) == 0 {
			return
		}

		slices.SortFunc(applicableEnemies, func(a Unit, b Unit) int {
			dur := GetEffectLongestRemainingDuration(a, Burn) - GetEffectLongestRemainingDuration(b, Burn)
			if dur != 0 {
				return dur
			}

			d1 := combat.GetLogicalDistance(unit.Position(), a.Position())
			d2 := combat.GetLogicalDistance(unit.Position(), b.Position())
			dist := d1 - d2
			if dist != 0 {
				return dist
			}
			return RandCmp()
		})

		cooldown += maximumCooldown
		target := applicableEnemies[0]

		unit.DealEffect(
			EffectBuilder(Burn).
				Source(unit).
				Target(target).
				Duration(capeDuration).
				Amount(burnAmount))

		unit.DealEffect(
			EffectBuilder(Wound).
				Source(unit).
				Target(target).
				Duration(capeDuration).
				Amount(woundAmount))

		LogC(combat, fmt.Sprintf("Sunfire Cape: %s->%s", unit.Name(), target.Name()))
	})
}

func RegisterThiefsGloves(unit Unit, isRadiant bool) {
	baseHealth := 150.0
	baseCrit := .20
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(baseHealth))
	unit.AddModifier(AdditiveModifier(CriticalStrikeChance).WithConstant(baseCrit))

	craftable := GetCraftableItems()

	// If we roll Thief's Gloves in RegisterThiefsGloves bad recursion will happen.
	// TODO: make following a function or something
	craftableNoThiefsGloves := make([]ItemName, 0, len(craftable))
	for _, item := range craftable {
		if item == ThiefsGloves {
			continue
		}
		craftableNoThiefsGloves = append(craftableNoThiefsGloves, item)
	}

	base := GetBaseItems()

	itemCraftable := craftableNoThiefsGloves[rand.IntN(len(craftable))]

	itemBase := base[rand.IntN(len(base))]

	// TODO: Find a way to display unit's rolled items
	RegisterItem(itemCraftable, false, unit)
	RegisterItem(itemBase, false, unit)

	rollMessage := fmt.Sprintf("%s rolled %s and %s", unit.Name(), ItemToString(itemCraftable), ItemToString(itemBase))
	LogC(unit.Combat(), rollMessage)
}

func RegisterTitansResolve(unit Unit, isRadiant bool) {

	baseAttackSpeed := 0.10
	baseArmorResist := 20.0
	unit.AddModifier(MultiplicativeModifier(AttackSpeed).WithConstant(baseAttackSpeed))
	unit.AddModifier(AdditiveModifier(ArmorResist).WithConstant(baseArmorResist))

	attackDamageStackMultiplier := 0.02
	maximumStacks := 25

	fullStacksArmorResist := 20.0
	fullStacksMagicResist := 20.0

	stacks := 0

	unit.RegisterEvent(OnAttack, func(event *Event) {
		stacks = MinInt(maximumStacks, stacks+1)
	})

	unit.RegisterEvent(OnAttacked, func(event *Event) {
		stacks = MinInt(maximumStacks, stacks+1)
	})

	unit.AddModifier(MultiplicativeModifier(AttackDamage).WithRetriever(
		func(unit Unit) float64 {
			return float64(stacks) * attackDamageStackMultiplier
		}))

	unit.AddModifier(AdditiveModifier(ArmorResist).WithRetriever(func(u Unit) float64 {
		if stacks >= maximumStacks {
			return fullStacksArmorResist
		}
		return 0
	}))
	unit.AddModifier(AdditiveModifier(MagicResist).WithRetriever(func(u Unit) float64 {
		if stacks >= maximumStacks {
			return fullStacksMagicResist
		}
		return 0
	}))
}

func RegisterWarmogsArmor(unit Unit, isRadiant bool) {
	baseHealthAdd := 600.0
	baseHealthMult := 0.08
	unit.AddModifier(AdditiveModifier(MaximumHealth).WithConstant(baseHealthAdd))
	unit.AddModifier(MultiplicativeModifier(MaximumHealth).WithConstant(baseHealthMult))
}

func SliceContains(slice []int, value int) bool {
	for _, k := range slice {
		if k == value {
			return true
		}
	}
	return false
}

func HealthPercentage(unit Unit, percentage float64) int {
	return RoundUpOptimistically(float64(unit.MaxHealth()) * percentage)
}
