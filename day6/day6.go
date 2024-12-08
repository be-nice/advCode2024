package day6

import (
	"fmt"
	"sync"
)

type Point struct {
	x, y int
}

type State struct {
	Point
	direction int
}

var directions = []Point{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

func turnRight(currentDirection int) int {
	return (currentDirection + 1) % 4
}

func Day6(r [][]rune) {
	var start Point
	var direction int

	for y, row := range r {
		found := false
		if found {
			break
		}
		for x, ch := range row {
			if ch == '^' || ch == '>' || ch == 'v' || ch == '<' {
				start = Point{x, y}
				switch ch {
				case '^':
					direction = 0
				case '>':
					direction = 1
				case 'v':
					direction = 2
				case '<':
					direction = 3
				}
				r[y][x] = '.'
				found = true
				break
			}
		}
	}

	s := State{start, direction}
	visited, _ := mxWalk(r, s, false)

	fmt.Println("Part 1")
	fmt.Println(len(visited))
	fmt.Println("Part 2")
	fmt.Println(part2(r, visited, s))
}

func mxWalk(r [][]rune, s State, loop bool) (map[Point]bool, bool) {
	visitedPoints := make(map[Point]bool)
	visitedStates := make(map[State]bool)
	pos := s.Point
	dir := s.direction

	for {
		if loop {
			state := State{pos, dir}

			if visitedStates[state] {
				return nil, true
			}

			visitedStates[state] = true
		} else {
			visitedPoints[pos] = true
		}

		nextPos := Point{pos.x + directions[dir].x, pos.y + directions[dir].y}

		if nextPos.y < 0 || nextPos.y >= len(r) || nextPos.x < 0 || nextPos.x >= len(r[0]) {
			return visitedPoints, false
		}

		if r[nextPos.y][nextPos.x] == '#' {
			dir = turnRight(dir)
		} else {
			pos = nextPos
		}
	}
}

func part2(r [][]rune, visited map[Point]bool, s State) int {
	loopingPositions := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for key := range visited {
		if key == s.Point {
			continue
		}

		wg.Add(1)
		go func(key Point) {
			defer wg.Done()
			cpR := deepCopy(r)
			cpR[key.y][key.x] = '#'

			_, loopDetected := mxWalk(cpR, s, true)

			if loopDetected {
				mu.Lock()
				loopingPositions++
				mu.Unlock()
			}
		}(key)
	}
	wg.Wait()

	return loopingPositions
}

func deepCopy(r [][]rune) [][]rune {
	cpR := make([][]rune, len(r))
	for i := range r {
		cpR[i] = make([]rune, len(r[i]))
		copy(cpR[i], r[i])
	}

	return cpR
}
