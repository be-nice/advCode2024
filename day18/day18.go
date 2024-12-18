package day18

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

const gridSize = 71

var (
	limit = 1024
	start = 0
)

var dirs = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

type Point struct {
	x, y int
}

func Day18(s []string) {
	grid := make([][]bool, gridSize)

	for i := range grid {
		grid[i] = make([]bool, gridSize)
	}

	parse(s, grid, map[Point]struct{}{})

	steps, path := shortestPath(grid)

	fmt.Println("Part 1")
	fmt.Println(steps)
	fmt.Println("Part 2")
	fmt.Println(part2(s, grid, path))
}

func part2(s []string, grid [][]bool, path map[Point]struct{}) string {
	start = limit + 1
	limit = len(s)
	steps := 0

	for {
		//fmt.Print("\n\n")
		//for i := range grid {
		//	for j := range grid[i] {
		//		if grid[j][i] {
		//			fmt.Print("#")
		//		} else if _, ok := path[Point{i, j}]; ok {
		//			fmt.Print("O")
		//		} else {
		//			fmt.Print(".")
		//		}
		//	}
		//	fmt.Print("\n")
		//}

		parse(s, grid, path)
		steps, path = shortestPath(grid)
		if steps == -1 {
			return s[start]
		}
		start++
	}
}

func parse(s []string, grid [][]bool, path map[Point]struct{}) {
	for ; start < limit; start++ {
		nxy := strings.Split(s[start], ",")
		px, _ := strconv.Atoi(nxy[0])
		py, _ := strconv.Atoi(nxy[1])

		grid[px][py] = true

		if _, ok := path[Point{px, py}]; ok {
			return
		}
	}

	return
}

func isValid(x, y int, grid [][]bool) bool {
	return x >= 0 && x < gridSize && y >= 0 && y < gridSize && !grid[y][x]
}

func shortestPath(grid [][]bool) (int, map[Point]struct{}) {
	queue := list.New()
	queue.PushBack(Point{0, 0})
	visited := make([][]bool, gridSize)

	for i := range visited {
		visited[i] = make([]bool, gridSize)
	}

	visited[0][0] = true

	path := make(map[Point]Point)
	steps := 0

	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			current := queue.Remove(queue.Front()).(Point)

			if current.x == gridSize-1 && current.y == gridSize-1 {
				resultPath := make(map[Point]struct{})

				for p := current; p != (Point{0, 0}); p = path[p] {
					resultPath[p] = struct{}{}
				}

				return steps, resultPath
			}

			for _, dir := range dirs {
				nx, ny := current.x+dir[0], current.y+dir[1]

				if isValid(nx, ny, grid) && !visited[ny][nx] {
					visited[ny][nx] = true
					nextPoint := Point{nx, ny}
					path[nextPoint] = current
					queue.PushBack(nextPoint)
				}
			}
		}

		steps++
	}

	return -1, nil
}
