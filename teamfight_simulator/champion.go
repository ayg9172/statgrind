package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Champion struct {
	Name               string
	Role               string
	Position           string
	Cost               int
	HexRange           int
	Health             []float64
	StartMana          float64
	Mana               float64
	AttackDamage       []float64
	Armor              float64
	MagicResist        float64
	AttackSpeed        float64
	Traits             []string
	AbilityPercentages []float64 // TODO: Implement AbilityPercentages

}

func inputString(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func inputFloat64(prompt string) float64 {
	var out float64
	fmt.Print(prompt)
	fmt.Scanln(&out)
	return out
}

func inputInt(prompt string) int {
	var out int
	fmt.Print(prompt)
	fmt.Scanln(&out)
	return out
}

// TODO: Use this combatreport to
// supplement abilities that require
// using amt of damage dealth/blocked over
// a period of time. Save a copy of a previous
// combat report

func enterLevels(statArray []float64, prompt string, levelCount int) []float64 {
	for i := 0; i < levelCount; i++ {
		fmt.Print(i + 1)
		statArray = append(statArray, inputFloat64(prompt))
	}
	return statArray
}

func _inputChampion() *Champion {
	champion := new(Champion)

	champion.Name = inputString("Champion Name: ")
	champion.Role = inputString("Role: ")
	champion.Position = inputString("Position: ")
	champion.Cost = inputInt("Cost: ")
	champion.HexRange = inputInt("Hex Range: ")

	champion.StartMana = inputFloat64("StartMana:")
	champion.Mana = inputFloat64("Mana:")

	champion.Health = make([]float64, 0, 3)
	champion.AttackDamage = make([]float64, 0, 3)
	champion.Health = enterLevels(champion.Health, "Health: ", 3)
	champion.AttackDamage = enterLevels(champion.AttackDamage, "AttackDamage: ", 3)

	champion.Armor = inputFloat64("Armor: ")
	champion.MagicResist = inputFloat64("Magic Resist: ")
	champion.AttackSpeed = inputFloat64("Attack Speed: ")

	champion.Traits = make([]string, 0, 3)

	trait := inputString("Enter trait: ")
	for trait != "" {
		champion.Traits = append(champion.Traits, trait)
		trait = inputString("Enter trait: ")
	}

	return champion
}

func LoadChampions() map[string]*Champion {
	files, _ := os.ReadDir("data")
	champs := make(map[string]*Champion)

	for _, value := range files {
		bytes, _ := os.ReadFile("data/" + value.Name())

		var champion *Champion

		json.Unmarshal(bytes, &champion)

		champs[champion.Name] = champion
	}
	return champs
}

func RoundDown(v float64) int {
	return int(v)
}

func RoundUpOptimistically(v float64) int {
	return int(v) + 1
}

func Resistance(armor int, damage int) int {
	a := float64(armor)
	d := float64(damage)
	return RoundDown(d * a / (a + 100))
}

func PercentageOf(damage int, percentage float64) int {
	return RoundDown(float64(damage) * percentage)
}

func Less(x int) func(int) bool {
	return func(y int) bool {
		return y < x
	}
}

func Equal(x int) func(int) bool {
	return func(y int) bool {
		return y == x
	}
}

func NotEqual(x int) func(int) bool {
	// TODO: Document that
	// NotEqual(0) filters minions, wolves, etc.
	return func(y int) bool {
		return y != x
	}
}

func Role(role string) func(*Champion) bool {
	return func(c *Champion) bool {
		return c.Role == role
	}
}
func createChampionJson() {
	champions := LoadChampions()
	jChamps, _ := json.Marshal(champions)
	os.WriteFile("data/champions.json", jChamps, 0644)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
