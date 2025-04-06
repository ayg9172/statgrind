package main

import "slices"

const noLimit = 0
const TeamOne TeamId = 0
const TeamTwo TeamId = 1
const Obstacle TeamId = 9

func OtherTeam(team TeamId) TeamId {
	if team == TeamOne {
		return TeamTwo
	}
	return TeamOne
}

func Order(units []Unit, ordering func(Unit, Unit) int) []Unit {
	slices.SortFunc(units, ordering)
	return units
}

/*********** [Filter] ***********/
func Filter(units []Unit, filters ...func(Unit) bool) []Unit {

	out := make([]Unit, 0, len(units))

	for _, unit := range units {
		includeUnit := true
		for _, filter := range filters {
			if !filter(unit) {
				includeUnit = false
				break
			}
		}
		if includeUnit {
			out = append(out, unit)

		}
	}

	return out
}

func IsAlive(unit Unit) bool {
	return unit.IsAlive()
}

func All() func(Unit) bool {

	return func(u Unit) bool {
		return true
	}
}

func SameTeam(unit Unit) func(Unit) bool {

	return func(other Unit) bool {
		return unit.Team() == other.Team()
	}
}

func DifferentTeam(unit Unit) func(Unit) bool {
	return func(other Unit) bool {
		return unit.Team() != other.Team()
	}
}

func adjacent(unit Unit) func(Unit) bool {
	return func(other Unit) bool {
		neighbors := unit.Combat().GetNeighbors(unit.Position())
		for _, i := range neighbors {
			if i == other.Position() {
				return true
			}
		}
		return false
	}
}

func WithinDistance(unit Unit, distance int) func(Unit) bool {
	if distance == 1 {
		return adjacent(unit)
	}
	return func(other Unit) bool {
		return unit.Combat().GetLogicalDistance(unit.Position(), other.Position()) <= distance
	}
}

func Affected(effect EffectName) func(Unit) bool {
	return func(unit Unit) bool {
		return unit.ContainsEffect(effect)
	}
}

func Not(filter func(Unit) bool) func(Unit) bool {
	return func(u Unit) bool {
		return !filter(u)
	}
}

/*********** [Orderings] ***********/
func ByAny() func(Unit, Unit) int {
	return func(_, _ Unit) int {
		return 0
	}
}

func ByRandom() func(Unit, Unit) int {
	return func(_, _ Unit) int {
		return RandCmp()
	}
}

func ByLogicalDistance(unit Unit) func(Unit, Unit) int {
	return Randomify(func(otherA Unit, otherB Unit) int {
		combat := unit.Combat()
		distanceA := combat.GetLogicalDistance(unit.Position(), otherA.Position())
		distanceB := combat.GetLogicalDistance(unit.Position(), otherB.Position())
		return distanceA - distanceB
	})
}

func ByMissingHealth() func(Unit, Unit) int {
	return Randomify(func(u1, u2 Unit) int {
		return u1.MissingHealth() - u2.MissingHealth()
	})
}

func ByHealthPercentage() func(Unit, Unit) int {
	return Randomify(func(u1, u2 Unit) int {
		percentage1 := float64(u1.CurrentHealth()) / float64(u1.MaxHealth())
		percentage2 := float64(u2.CurrentHealth()) / float64(u2.MaxHealth())

		if percentage1 <= percentage2 {
			return -1
		}
		return 1
	})
}

func Reverse(ordering func(Unit, Unit) int) func(Unit, Unit) int {
	return func(u1, u2 Unit) int {
		return -ordering(u1, u2)
	}
}

func Randomify(ordering func(Unit, Unit) int) func(Unit, Unit) int {
	return func(u1, u2 Unit) int {
		diff := ordering(u1, u2)
		if diff == 0 {
			return RandCmp()
		}
		return diff
	}
}

/************ [Search Options] ************/

type SearchOptions struct {
	source  Unit
	limit   int
	cmp     func(Unit, Unit) int
	filters []func(Unit) bool
}

func newSearchOptions(unit Unit) *SearchOptions {
	return &SearchOptions{
		source:  unit,
		limit:   noLimit,
		cmp:     ByAny(),
		filters: make([]func(Unit) bool, 0, 3),
	}
}

func Friends(unit Unit) *SearchOptions {
	out := newSearchOptions(unit)
	out.filters = append(out.filters, SameTeam(unit))
	return out
}

func Enemies(unit Unit) *SearchOptions {
	out := newSearchOptions(unit)
	out.filters = append(out.filters, DifferentTeam(unit))
	return out
}

func (options *SearchOptions) WithLimit(limit int) *SearchOptions {
	options.limit = limit
	return options
}

func (options *SearchOptions) WithDistanceLimit(distance int) *SearchOptions {
	options.filters = append(options.filters, WithinDistance(options.source, distance))
	return options
}

func (options *SearchOptions) WithByRandom() *SearchOptions {
	options.cmp = ByRandom()
	return options
}

func (options *SearchOptions) WithByDistance() *SearchOptions {
	options.cmp = ByLogicalDistance(options.source)
	return options
}

func (options *SearchOptions) WithReverse() *SearchOptions {
	options.cmp = Reverse(options.cmp)
	return options
}
