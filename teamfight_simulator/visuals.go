package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const HealthFactor = 100

const (
	Reset     = "\033[0m"
	Bold      = "\033[1m"
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	LightBlue = "\033[38;5;123m"
	Grey      = "\033[90m"

	Underline        = "\033[4m"
	Cyan             = "\033[36m"
	White            = "\033[37m"
	LightRed         = "\033[91m"
	BoldBlack        = "\033[1;30m"
	BoldRed          = "\033[1;31m"
	BoldGreen        = "\033[1;32m"
	BoldYellow       = "\033[1;33m"
	BoldBlue         = "\033[1;34m"
	BoldMagenta      = "\033[1;35m"
	BoldCyan         = "\033[1;36m"
	BoldWhite        = "\033[1;37m"
	RedBackground    = "\033[41m"
	GreenBackground  = "\033[42m"
	YellowBackground = "\033[43m"
	BlueBackground   = "\033[44m"
	PurpleBackground = "\033[45m"
	CyanBackground   = "\033[46m"
	WhiteBackground  = "\033[47m"
)

func sf64(f64 float64) string {
	out := fmt.Sprintf("%.2f", f64)
	float2f, _ := strconv.ParseFloat(out, 64)
	if float2f == float64(int(float2f)) {
		out = fmt.Sprintf("%.0f", f64)
	}

	return out
}

func hex(s string) string {
	if s == "" {
		s = "_"
	}

	return "\\" + s + Reset + "/"
}

const RightArrow = "-> "

const UNDERLINE_SEP = true

func getSeparator(index int, symbols []string, getNeighborsFunction func(int) []int) string {

	mapping := getIndexStringMapping()
	neighbors := getNeighborsFunction(mapping[index])

	mIndex := mapping[index]
	row := mIndex / 7

	mNeighborIndex := mIndex - 7
	if row%2 == 0 {
		mNeighborIndex = mIndex - 8
	}

	if IsPresent(neighbors, mNeighborIndex) {
		if symbols[mNeighborIndex] != "" {
			return "⎺"
		}
	}
	return " "

}

func board(symbols []string) string {
	var b strings.Builder
	mapping := getIndexStringMapping()

	// TODO: This line cane be optimized for higher FPS
	gN := CreateGetNeighborsFunction(8, 7)

	b.WriteString("  ")
	b.WriteString(strings.Repeat(" ", len(RightArrow)))
	b.WriteString("/‾\\_/‾\\_/‾\\_/‾\\")
	b.WriteString("\n" + RightArrow + "/‾")

	index := 0
	for r := 0; r < 7; r++ {
		for i := 0; i < 4; i++ {

			placement := symbols[mapping[index]]
			b.WriteString(hex(placement) + getSeparator(index, symbols, gN))
			index += 1

		}
		b.WriteString("\n" + RightArrow)
		for i := 0; i < 4; i++ {

			if i == 3 && r != 6 {
				placement := symbols[mapping[index]]
				b.WriteString(hex(placement) + " \\\n" + RightArrow + "/‾")
				index += 1
				break
			}

			placement := symbols[mapping[index]]
			b.WriteString(hex(placement) + getSeparator(index, symbols, gN))
			index += 1
		}

	}

	return b.String()
}

func PopulateBoard(units []Unit) string {
	b := make([]string, 8*7)

	for _, unit := range units {
		if unit == nil || unit.Position() < 0 {
			continue
		}
		color := BoldGreen
		if unit.Team() == TeamTwo {
			color = Red
		}
		b[unit.Position()] = color + string(unit.Name()[0]) + Reset
	}

	return board(b)
}

func healthbar(health int, maxHealth int, shield int, color string) string {

	if health > maxHealth {
		panic("Health>MaxHealth")
	}
	if health <= 0 {
		return "[ Eliminated ]"
	}

	var b strings.Builder

	spaceF := float64(maxHealth) / HealthFactor
	barsF := float64(health) / HealthFactor
	shieldBarsF := float64(shield) / HealthFactor

	maxHealth += shield

	space := int(math.Round(spaceF))
	bars := int(math.Round(barsF))
	shieldBars := int(math.Round(shieldBarsF))

	b.WriteString(BoldWhite)
	b.WriteString("[")
	b.WriteString(color)
	b.WriteString(strings.Repeat("■", bars))
	b.WriteString(BoldWhite)
	b.WriteString(strings.Repeat("■", shieldBars))

	b.WriteString(Reset)
	b.WriteString(strings.Repeat(" ", space-bars))
	b.WriteString(BoldWhite)
	b.WriteString("]")
	b.WriteString(Reset)

	return b.String()
}

func ClearTerminal() {
	fmt.Print("\033c")
	/*
		if runtime.GOOS == "windows" {
			cmd := exec.Command("cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		} else {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}
	*/
}

const COMBAT_HISTORY_LENGTH = 30

func Display(combat *Combat) {
	champions := DisplayUnits(combat.Units, combat)
	board := PopulateBoard(combat.Units)
	bufferedWriter := bufio.NewWriter(os.Stdout)

	// TODO: Ideally this clearterminal is right before the flush
	// TODO: So that we're not thinking while terminal is empty
	// TODO: However, this results in champ stats and hex board being cut off / empty
	// TODO: So for now this clear terminal is here
	ClearTerminal()

	fmt.Fprintf(bufferedWriter, "%s\n%s", champions, board)

	fmt.Fprintf(bufferedWriter, "\n\n========== Combat Log ==========\n")

	for i := 0; i < MinInt(COMBAT_HISTORY_LENGTH, len(combat.CombatLog)); i++ {
		fmt.Fprintf(bufferedWriter, "\n%s", combat.CombatLog[len(combat.CombatLog)-i-1])
	}

	for i := 0; i < COMBAT_HISTORY_LENGTH-len(combat.CombatLog); i++ {
		fmt.Fprintln(bufferedWriter)
	}
	fmt.Fprintf(bufferedWriter, "\n================================\n")

	bufferedWriter.Flush()
}

func GetUnitNameLess(units []Unit) func(x int, y int) bool {
	return func(x int, y int) bool {

		comparison := cmp.Compare(units[x].Name(), units[y].Name())
		if comparison == 0 {
			return units[x].Id() < units[y].Id()
		}

		return comparison < 0
	}
}

func DisplayUnits(units []Unit, combat *Combat) string {
	var out strings.Builder
	team1 := Filter(units, func(u Unit) bool { return u.Team() == TeamOne })
	team2 := Filter(units, func(u Unit) bool { return u.Team() == TeamTwo })
	sort.Slice(team1, GetUnitNameLess(team1))
	sort.Slice(team2, GetUnitNameLess(team2))

	for _, c := range team2 {
		out.WriteString(DisplayUnit(c) + "\n")
	}
	out.WriteString("\n")
	for _, c := range team1 {
		out.WriteString(DisplayUnit(c) + "\n")
	}
	return out.String()
}

// ensure zero
func ez(val int) int {
	if val < 0 {
		return 0
	}
	return val
}

func DisplayUnit(unit Unit) string {
	name := unit.Name()
	health := unit.CurrentHealth()

	shieldSum, _ := GetActiveEffectSum(unit, Shield)
	shield := int(shieldSum)
	maxHealth := unit.MaxHealth()
	abilityPower := unit.AbilityPower()
	mana := unit.CurrentMana()
	maxMana := unit.MaxMana()
	attackDamage := unit.AttackDamage()
	damageAmp := unit.DamageAmp()
	durability := unit.Durability()
	omnivamp := unit.Omnivamp()
	criticalChance := unit.CriticalStrikeChance()
	criticalDamage := unit.CriticalStrikeDamage()
	attackSpeed := unit.AttackSpeed()
	magicResist := unit.MagicResist()
	armorResist := unit.ArmorResist()

	var b strings.Builder

	nameColor := Red
	if unit.Team() == TeamOne {
		nameColor = BoldGreen
	}

	b.WriteString(nameColor + name + Reset)

	if unit.Level() == 0 {
		b.WriteString(" 1*" + Reset)
	}
	if unit.Level() == 1 {
		b.WriteString(Bold + " 2*" + Reset)
	}
	if unit.Level() == 2 {
		b.WriteString(BoldYellow + " 3*" + Reset)
	}

	b.WriteString(strings.Repeat(" ", ez(15-len(name))))

	b.WriteString(healthbar(health, maxHealth, shield, nameColor))

	if health > 0 {
		// Display mana bar
		b.WriteString(BoldWhite + "[" + Reset + LightBlue + strconv.Itoa(unit.CurrentMana()))
		b.WriteString(Reset + "/")
		b.WriteString(LightBlue + strconv.Itoa(unit.CurrentMana()))
		b.WriteString(BoldWhite + "]" + Reset)
	}

	hp := int(math.Round(float64(maxHealth)/HealthFactor)) + 2
	mb := len(strconv.Itoa(mana)) + len(strconv.Itoa(maxMana)) + 3

	if health <= 0 {
		hp = len("[ Eliminated ]")
		mb = 0
	}

	b.WriteString(strings.Repeat(" ", ez(60-hp-mb)))

	ad := len(strconv.Itoa(attackDamage))
	b.WriteString(Yellow + strconv.Itoa(attackDamage))

	b.WriteString(strings.Repeat(" ", ez(4-ad)))

	as := len(sf64(attackSpeed))

	b.WriteString(Reset + "*")
	b.WriteString(Yellow + sf64(attackSpeed))

	b.WriteString(strings.Repeat(" ", ez(4-as)))

	b.WriteString(Reset + " | ")
	ap := len(strconv.Itoa(abilityPower))

	b.WriteString(Cyan + strconv.Itoa(abilityPower))
	b.WriteString(strings.Repeat(" ", ez(10-ap)))

	b.WriteString(BoldYellow + strconv.Itoa(armorResist) + Reset)

	b.WriteString(strings.Repeat(" ", ez(4-len(strconv.Itoa(armorResist)))))

	b.WriteString("| ")
	b.WriteString(BoldCyan + strconv.Itoa(magicResist) + Reset)

	b.WriteString(strings.Repeat(" ", ez(6-len(strconv.Itoa(magicResist)))))

	b.WriteString("    ")
	b.WriteString(BoldRed + sf64(criticalChance*100) + "% ")
	b.WriteString(BoldRed + sf64(criticalDamage*100) + "% " + Reset)
	b.WriteString("    ")

	if damageAmp > 0 {
		b.WriteString(Reset + "Amp:" + sf64(damageAmp*100) + "% ")
	}

	if durability > 0 {
		b.WriteString(Reset + "Dur:" + sf64(durability*100) + "%")
	}

	if omnivamp > 0 {
		b.WriteString(Reset + "Omni:" + sf64(omnivamp*100) + "% ")
	}

	return b.String()
}
