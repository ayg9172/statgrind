package main

import (
	"math"
	"math/rand"
	"slices"
)

const FRONTLINE_ALLY = 4
const FRONTLINE_ENEMY = 3

func IsPresent(arr []int, i int) bool {
	for _, k := range arr {
		if k == i {
			return true
		}
	}
	return false
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxInt(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func MinInt(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func CreateGetLogicalDistanceFunction8x7() func(int, int) int {
	dist := make([][]int, 8*7)

	for i := range dist {
		dist[i] = make([]int, 8*7)
	}

	for x := 0; x < 7*8; x++ {
		distances, _ := BFS(x, make([]Unit, 8*7), TeamOne, CreateGetNeighborsFunction(8, 7))

		for y := 0; y < 7*8; y++ {
			dist[x][y] = distances[y]
		}
	}
	return func(start int, end int) int {
		return dist[start][end]
	}
}

func dist(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	h := x2 - x1
	v := y2 - y1
	return math.Sqrt(h*h + v*v)
}

func PhysicalDist(hex1 int, hex2 int) float64 {
	// TODO: unit test this
	x1, y1 := GetAbsoluteCoordinate8x7(hex1)
	x2, y2 := GetAbsoluteCoordinate8x7(hex2)
	return dist(x1, y1, x2, y2)
}

func getHexIndex8x6(row int, col int) int {
	return row*7 + col
}

func GetAbsoluteCoordinate8x7(i int) (float64, float64) {
	row := i / 7
	col := i % 7

	x := 0.5
	y := 0.5

	if row%2 == 1 {
		x = 1
	}

	return y + float64(row), x + float64(col)
}

func MakeReversePath(start int, end int, previous []int) []int {
	path := make([]int, 0, 10)
	current := end
	for current != start {
		path = append(path, current)
		current = previous[current]

	}
	return append(path, start)
}

func MakePath(start int, end int, previous []int) []int {
	reversePath := MakeReversePath(start, end, previous)
	slices.Reverse(reversePath)
	return reversePath
}

func FindPath(start int, end int, board []Unit, team TeamId, getNeighbors func(int) []int) []int {

	if start < 0 || end < 0 {
		return make([]int, 0)
	}

	queue := NewQueue()
	previous := make([]int, len(board))

	for i := range previous {
		previous[i] = -1
	}

	previous[end] = end

	queue.Enqueue(end)

	for queue.Size() > 0 {

		c := queue.Pop()
		n := getNeighbors(c)

		for _, neighbor := range n {

			if neighbor == start {
				previous[start] = c
				return MakeReversePath(end, start, previous)
			}

			if previous[neighbor] >= 0 {
				continue
			}

			if board[neighbor] != nil && team == board[neighbor].Team() {
				continue
			}

			previous[neighbor] = c
			queue.Enqueue(neighbor)

		}

	}

	return make([]int, 0)
}

func BFS(start int, board []Unit, team TeamId, getNeighbors func(int) []int) ([]int, []int) {

	if start < 0 {
		return make([]int, 0), make([]int, 0)
	}

	queue := NewQueue()
	distanceQueue := NewQueue()

	queue.Enqueue(start)
	distanceQueue.Enqueue(0)

	previous := make([]int, len(board))
	distances := make([]int, len(board))
	for i := range distances {
		distances[i] = -1
	}
	for i := range previous {
		previous[i] = -1
	}

	previous[start] = start

	for queue.Size() > 0 {

		currentIndex := queue.Pop()
		currentDistance := distanceQueue.Pop()

		distances[currentIndex] = currentDistance

		n := getNeighbors(currentIndex)

		for _, neighbor := range n {

			if previous[neighbor] >= 0 {
				continue
			}

			if board[neighbor] != nil && team == board[neighbor].Team() {
				continue
			}

			previous[neighbor] = currentIndex
			queue.Enqueue(neighbor)
			distanceQueue.Enqueue(currentDistance + 1)
		}
	}

	return distances, previous
}

type Compass struct {
	Northwest func(int) int
	Northeast func(int) int
	Southwest func(int) int
	Southeast func(int) int
	East      func(int) int
	West      func(int) int
}

func Northwest(index int) int {
	row := index / 7
	col := index % 7

	isEvenRow := row%2 == 0

	if row == 0 || (col == 0 && isEvenRow) {
		return -1
	}

	if isEvenRow {
		return index - 8
	} else {
		return index - 7
	}
}

func Northeast(index int) int {
	row := index / 7
	col := index % 7
	isEvenRow := row%2 == 0
	if row == 0 || (col == 6 && !isEvenRow) {
		return -1
	}

	if isEvenRow {
		return index - 7
	} else {
		return index - 6
	}
}

func Southwest(index int) int {
	row := index / 7
	col := index % 7
	isEvenRow := row%2 == 0
	if row == 7 || (col == 0 && isEvenRow) {
		return -1
	}
	if isEvenRow {
		return index + 6
	} else {
		return index + 7
	}
}

func Southeast(index int) int {
	row := index / 7
	col := index % 7
	isEvenRow := row%2 == 0
	if row == 7 || (col == 6 && !isEvenRow) {
		return -1
	}
	if isEvenRow {
		return index + 7
	} else {
		return index + 8
	}
}

func East(index int) int {
	col := index % 7
	if col == 6 {
		return -1
	}
	return index + 1
}

func West(index int) int {
	col := index % 7
	if col == 0 {
		return -1
	}
	return index - 1
}

// TODO: Document this functionality
func CreateCircleSearchTargetSort() func(TeamId, int, []Unit) {
	getDist := CreateGetLogicalDistanceFunction8x7()
	return func(team TeamId, origin int, potentialTargets []Unit) {
		if len(potentialTargets) == 0 {
			return
		}

		distance := getDist(origin, potentialTargets[0].Position())

		// This matters if distance is odd
		forwardDirectionOrder := rand.Intn(2)

		for i := 0; i < distance; i++ {
			if team == TeamOne {
				if i%2 == forwardDirectionOrder {
					origin = Northeast(origin)
				} else {
					origin = Northwest(origin)
				}
			} else {
				if i%2 == forwardDirectionOrder {
					origin = Southwest(origin)
				} else {
					origin = Southeast(origin)
				}
			}
		}

		slices.SortFunc(potentialTargets, func(a Unit, b Unit) int {
			distA := getDist(origin, a.Position())
			distB := getDist(origin, b.Position())

			// TODO: Document this random target selection
			if distA == distB {
				return RandCmp()
			}

			return distA - distB
		})
	}
}

func RandCmp() int {
	if rand.Intn(2) == 1 {
		return 1
	}
	return -1
}

func CreateTargetSelector() func(int, *Combat, TeamId) ([]Unit, []int) {
	circleSearchTargetSort := CreateCircleSearchTargetSort()
	getNeighbors := CreateRandomGetNeighborsFunc()
	// TODO: This should return a bool ok as second return if new target was found
	return func(start int, combat *Combat, team TeamId) ([]Unit, []int) {
		// TODO: Use unit for the start param? Actually
		// Document we use int for start, because we may want to search
		// From non 0 coordinates
		// ALSO, ADD OPTION TO EXCLUDE TARGET DUMMIES?
		// Depends on the behavior in the real game
		distances, previous := BFS(start, combat.Board, team, getNeighbors)

		enemyTeamFilter := func(unit Unit) bool { return unit.Team() == OtherTeam(team) }

		/** TODO: BIG TODO: IMPORTANT: This path finding works for melee units
		HOWEVER,  ranged units do not need to path to the unit, they need to path
		to the closest hex from where they can ranged attack the target **/

		// TODO: By this point any unit a ranged unit cannot reach, should have a dist -1
		// TODO: For ranged targets, we need to know the hex we will path towards!!

		enemyUnits := Filter(combat.Units, enemyTeamFilter)
		enemyUnits = Filter(enemyUnits, IsAlive)
		enemyUnits = Filter(enemyUnits, func(otherUnt Unit) bool {
			// TODO: Instead of u.Position, we use the hex we're pathing towards
			// Tis is to accomodate ranged units
			return distances[otherUnt.Position()] >= 0
		})

		if len(enemyUnits) == 0 {
			// TODO: Re-add this panic AFTER ADDING MORE ROBUST COMBAT FINISH METHOD
			// panic("Target Selection called when there are no remaining enemy units")
			return make([]Unit, 0), make([]int, 0)
		}

		slices.SortFunc(enemyUnits, func(x, y Unit) int {
			return distances[x.Position()] - distances[y.Position()]
		})

		minimumDistance := distances[enemyUnits[0].Position()]

		targets := make([]Unit, 0, len(combat.Units))
		for _, unit := range enemyUnits {

			if minimumDistance == distances[unit.Position()] {
				targets = append(targets, unit)
			}
		}

		circleSearchTargetSort(team, start, targets)
		return targets, previous
	}
}

func CreateGetNeighborsFunction(rows int, columns int) func(int) []int {
	adjacency := BuildAdjacencySlice(rows, columns)
	return func(x int) []int {
		return adjacency[x]
	}
}

func BuildAdjacencySlice(rows int, columns int) [][]int {
	size := rows * columns

	adjacency := make([][]int, size)
	for i := 0; i < size; i++ {

		neighbors := make([]int, 0, columns)

		if i%columns != columns-1 {
			neighbors = append(neighbors, i+1)
		}

		if i < size-columns {
			neighbors = append(neighbors, i+columns)

			if (i/7)%2 == 0 {
				if i%7 != 0 {
					// leftdown
					neighbors = append(neighbors, i+columns-1)
				}
			} else {
				if i%7 != 6 {
					// rightdown
					neighbors = append(neighbors, i+columns+1)

				}
			}
		}

		adjacency[i] = neighbors
	}

	for i, neighbors := range adjacency {

		for _, n := range neighbors {
			if !IsPresent(adjacency[n], i) {
				adjacency[n] = append(adjacency[n], i)
			}
		}
	}

	for i := range adjacency {
		slices.Sort(adjacency[i])
	}
	return adjacency
}

func getIndexStringMapping() []int {
	mapping := []int{
		42, 28, 14, 0, 49, 35, 21,
		7, 43, 29, 15, 1, 50, 36,
		22, 8, 44, 30, 16, 2, 51,
		37, 23, 9, 45, 31, 17, 3,
		52, 38, 24, 10, 46, 32, 18,
		4, 53, 39, 25, 11, 47, 33,
		19, 5, 54, 40, 26, 12, 48,
		34, 20, 6, 55, 41, 27, 13,
	}

	return mapping
}
