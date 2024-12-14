package day14

import (
	"fmt"
	"strconv"
	"strings"
)

type quadrant struct {
	q1, q2, q3, q4 int
}

type robot struct {
	pX, pY, vX, vY int
}

const (
	SimulateSeconds = 100
	MaxHeight       = 103
	MaxWidth        = 101
	MidHeight       = MaxHeight / 2
	MidWidth        = MaxWidth / 2
)

func (q *quadrant) getQuadrant(x, y int) {
	if x == MidWidth || y == MidHeight {
		return
	}

	if y < MidHeight {
		if x < MidHeight {
			q.q1++
		} else {
			q.q2++
		}
	} else {
		if x < MidHeight {
			q.q3++
		} else {
			q.q4++
		}
	}
}

func (r *robot) calcNewPos(offset int) {
	r.pX = (r.pX + r.vX*offset) % MaxWidth
	r.pY = (r.pY + r.vY*offset) % MaxHeight

	if r.pX < 0 {
		r.pX += MaxWidth
	}

	if r.pY < 0 {
		r.pY += MaxHeight
	}
}

func Day14(s []string) {
	q := &quadrant{}

	robots := parse(s)

	for _, robot := range robots {
		robot.calcNewPos(SimulateSeconds)
		q.getQuadrant(robot.pX, robot.pY)
	}

	// Matrix printing is handled in Part 2
	// calling it first leaves numeric output as last lines
	p2 := part2(robots)

	fmt.Println("Part 1")
	fmt.Println(q.q1 * q.q2 * q.q3 * q.q4)
	fmt.Println("Part 2")
	fmt.Println(p2)
}

func part2(robots []*robot) int {
	// Reusing robot locations from part 1
	// Starting from SimulateSeconds + 1 offset to have correct part 2 count
	for i := SimulateSeconds + 1; i < 10000; i++ {
		q := &quadrant{}

		for _, r := range robots {
			r.calcNewPos(1)
			q.getQuadrant(r.pX, r.pY)
		}

		// MAGIC NUMBERS
		// MAGIC NUMBERS
		// MAGIC NUMBERS

		if q.q1 > 205 || q.q2 > 205 || q.q3 > 205 || q.q4 > 205 {
			if buildMatrix(robots) {
				return i
			}
		}
	}

	return -1
}

func buildMatrix(robots []*robot) bool {
	matrix := make([][]rune, MaxHeight)

	for m := range matrix {
		matrix[m] = make([]rune, MaxWidth)

		for n := range matrix[m] {
			matrix[m][n] = '.'
		}
	}

	for _, r := range robots {
		matrix[r.pY][r.pX] = '#'
	}

	if checkPattern(matrix) {
		printMatrix(matrix)

		return true
	}

	return false
}

func checkPattern(matrix [][]rune) bool {
	pattern := []string{
		"..#..",
		".###.",
		"#####",
	}

	pHeight := len(pattern)
	pWidth := len(pattern[0])

	for y := 0; y <= len(matrix)-pHeight; y++ {
		for x := 0; x <= len(matrix[0])-pWidth; x++ {
			match := true
			for dy := 0; dy < pHeight; dy++ {
				for dx := 0; dx < pWidth; dx++ {
					if rune(pattern[dy][dx]) != matrix[y+dy][x+dx] {
						match = false
						break
					}
				}

				if !match {
					break
				}
			}

			if match {
				return true
			}
		}
	}

	return false
}

func printMatrix(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}

func parse(s []string) []*robot {
	robots := make([]*robot, 0, len(s))

	for _, line := range s {
		r := &robot{}
		tokens := strings.Fields(line)

		p := strings.Split(strings.TrimPrefix(tokens[0], "p="), ",")
		v := strings.Split(strings.TrimPrefix(tokens[1], "v="), ",")

		r.pX, _ = strconv.Atoi(p[0])
		r.pY, _ = strconv.Atoi(p[1])
		r.vX, _ = strconv.Atoi(v[0])
		r.vY, _ = strconv.Atoi(v[1])

		robots = append(robots, r)
	}

	return robots
}
