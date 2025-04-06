package main

type Player struct {
	Health   int
	Gold     int
	Streak   int
	Augments []string
	Items    []string
	Anomaly  string
}

func CreateDefaultPlayer() *Player {
	return &Player{
		Health:   100,
		Gold:     0,
		Streak:   0,
		Augments: make([]string, 3),
		Items:    make([]string, 10),
	}
}
