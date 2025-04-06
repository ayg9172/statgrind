package main

func testCombat(unit string, speed float64) {

	combat := CreateDefaultCombat()

	combat.NewUnit(unit, 0, getHexIndex8x6(6, 3), TeamOne)
	combat.Units[0].AddItem(SpearOfShojin)
	combat.Units[0].AddItem(Crownguard)
	combat.Units[0].AddItem(AdaptiveHelm)

	// combat.NewUnit(unit, 0, getHexIndex8x6(7, 3), true)

	// trundle := createUnit("Trundle", 0, getHexIndex8x6(7, 4), true)

	combat.NewUnit("Minion", 0, getHexIndex8x6(3, 1), TeamTwo)
	combat.NewUnit("Minion", 0, getHexIndex8x6(3, 3), TeamTwo)

	CombatSlowRun(combat, speed)
}

func main() {

	testCombat("Lux", 1)
	// testCombat2("Sevika", 5)
	// time.Sleep(time.Second * 2)
	// testCombat3("Sevika", 5)
	// runSimpleTargetTests()
	// testPathFinding()
}
