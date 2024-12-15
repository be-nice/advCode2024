package day15

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y int
}

func (p *Point) Add(dir Point) Point {
	return Point{p.X + dir.X, p.Y + dir.Y}
}

var dir = map[rune]Point{
	'^': {0, -1},
	'>': {1, 0},
	'v': {0, 1},
	'<': {-1, 0},
	'[': {1, 0},
	']': {-1, 0},
}

func Day15(s string) {
	split := strings.Split(strings.TrimSpace(s), "\n\n")
	r := strings.NewReplacer("#", "##", "O", "[]", ".", "..", "@", "@.")

	fmt.Println("Part 1")
	fmt.Println(solve(split[0], split[1]))
	fmt.Println("Part 2")
	fmt.Println(solve(r.Replace(split[0]), split[1]))
}

func solve(input, moves string) int {
	grid := make(map[Point]rune)
	robot := Point{}

	for y, line := range strings.Fields(input) {
		for x, ch := range line {
			if ch == '@' {
				robot = Point{x, y}
			}

			grid[Point{x, y}] = ch
		}
	}

loop:
	for _, ch := range strings.ReplaceAll(moves, "\n", "") {
		queue := []Point{robot}
		boxes := make(map[Point]rune)

		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]

			if _, ok := boxes[p]; ok {
				continue
			}

			boxes[p] = grid[p]

			switch n := p.Add(dir[ch]); grid[n] {
			case '#':
				continue loop
			case '[', ']':
				queue = append(queue, n.Add(dir[grid[n]]))
				fallthrough
			case 'O':
				queue = append(queue, n)
			}
		}

		for box := range boxes {
			grid[box] = '.'
		}
		for box := range boxes {
			grid[box.Add(dir[ch])] = boxes[box]
		}
		robot = robot.Add(dir[ch])
	}

	res := 0

	for p, val := range grid {
		if val == 'O' || val == '[' {
			res += 100*p.Y + p.X
		}
	}

	return res
}
