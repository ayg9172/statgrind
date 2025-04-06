package main

func getTraitRequirements(traitName string) []int {

	// TODO: This should be replaced with JSON File
	switch traitName {
	case "Academy":
		return []int{3, 4, 5, 6}
	case "Automata":
		return []int{2, 4, 6}
	case "BlackRose":
		return []int{3, 4, 5, 7}
	case "ChemBaron":
		return []int{3, 4, 5, 6, 7}
	case "Conqueror":
		return []int{2, 4, 6, 9}
	case "Emissary":
		return []int{1, 4}
	case "Enforcer":
		return []int{2, 4, 6, 8, 10}
	case "Experiment":
		return []int{3, 5, 7}
	case "Family":
		return []int{3, 4, 5}
	case "Firelight":
		return []int{2, 3, 4}
	case "High Roller":
		return []int{1}
	case "Junker King":
		return []int{1}
	case "Rebel":
		return []int{3, 5, 7, 10}
	case "Scrap":
		return []int{2, 4, 6, 9}
	case "Ambusher":
		return []int{2, 3, 4, 5}
	case "Artillerist":
		return []int{2, 4, 6}
	case "Bruiser":
		return []int{2, 4, 6}
	case "Dominator":
		return []int{2, 4, 6}
	case "Form Swapper":
		return []int{2, 4}
	case "Pit Fighter":
		return []int{2, 4, 6, 8}
	case "Quickstriker":
		return []int{2, 3, 4}
	case "Sentinel":
		return []int{2, 4, 6}
	case "Sniper":
		return []int{2, 4, 6}
	case "Sorcerer":
		return []int{2, 4, 6, 8}
	case "Visionary":
		return []int{2, 4, 6, 8}
	case "Watcher":
		return []int{2, 4, 6}

	default:
		panic("Unimplemented Trait for Levels: " + traitName)
	}
}

func GetTraitLevel(traitName string, traitCount int) int {
	requirements := getTraitRequirements(traitName)
	level := 0
	for _, requirement := range requirements {
		if traitCount >= requirement {
			level += 1
		} else {
			return level
		}
	}
	return level
}
