package main

import "math"

type Tick int

const CHAMPION_MOVEMENT_TICKS = 15

func TickRoundingScheme(tick float64) int {
	return int(math.Round(tick))
}

func SecondsToTicksRound(seconds int) int {
	return TickRoundingScheme(float64(seconds) * DefaultTickRate)
}

func SecondsToTicks(seconds float64) float64 {
	return seconds * DefaultTickRate
}

func CalculateAttackDurationTicks(attackSpeed float64) int {
	if attackSpeed == 0 {
		return SecondsToTicksRound(999999)
	}
	return TickRoundingScheme(DefaultTickRate / attackSpeed)
}
