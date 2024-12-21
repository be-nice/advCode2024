package day20

import (
	"fmt"
	"strings"
)

type Cell byte

const (
	Empty Cell = iota
	Wall
	Start
	End
)

type Direction struct {
	dx, dy int
}

type point struct {
	x, y int
}

type grid struct {
	data          [][]Cell
	width, height int
}

var (
	Up    = Direction{0, -1}
	Down  = Direction{0, 1}
	Left  = Direction{-1, 0}
	Right = Direction{1, 0}
)

func (d Direction) RotateRight() Direction {
	return Direction{-d.dy, d.dx}
}

func (d Direction) RotateLeft() Direction {
	return Direction{d.dy, -d.dx}
}

func (p point) add(d Direction) point {
	return point{p.x + d.dx, p.y + d.dy}
}

func NewGrid(width, height int) *grid {
	data := make([][]Cell, height)
	for i := range data {
		data[i] = make([]Cell, width)
	}

	return &grid{data: data, width: width, height: height}
}

func (g *grid) getPos(p point) Cell {
	if p.x < 0 || p.x >= g.width || p.y < 0 || p.y >= g.height {
		return Wall
	}

	return g.data[p.y][p.x]
}

func (g *grid) set(p point, c Cell) {
	if p.x >= 0 && p.x < g.width && p.y >= 0 && p.y < g.height {
		g.data[p.y][p.x] = c
	}
}

func (g *grid) find(c Cell) point {
	for y, row := range g.data {
		for x, cell := range row {
			if cell == c {
				return point{x, y}
			}
		}
	}

	return point{}
}

func Day20(b []byte) {
	grid := Parse(b)

	fmt.Println(findCheats(grid, 100, 2))
	fmt.Println(findCheats(grid, 100, 20))
}

func Parse(buf []byte) *grid {
	lines := strings.Split(strings.TrimSpace(string(buf)), "\n")
	height := len(lines)
	width := len(lines[0])
	grid := NewGrid(width, height)

	for y, line := range lines {
		for x, char := range line {
			var cell Cell
			switch char {
			case '.':
				cell = Empty
			case '#':
				cell = Wall
			case 'S':
				cell = Start
			case 'E':
				cell = End
			default:
				return nil
			}
			grid.set(point{x, y}, cell)
		}
	}

	return grid
}

func findCheats(g *grid, save, cheat int) int {
	start := g.find(Start)
	d := Up
	path := []point{start}
	p := start

	for g.getPos(p) != End {
		for _, dd := range []Direction{d, d.RotateRight(), d.RotateLeft()} {
			if q := p.add(dd); g.getPos(q) != Wall {
				path, p, d = append(path, q), q, dd
				break
			}
		}
	}

	count := 0
	for i, p := range path[:len(path)-save] {
		for j, q := range path[i+save:] {
			if d := l1(p, q); d <= cheat && d <= j {
				count++
			}
		}
	}

	return count
}

func l1(p1, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
