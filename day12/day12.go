package day12

import "fmt"

const SIZE = 140

type Point struct {
	x, y int
}

var (
	grid   [SIZE][SIZE]byte
	shapes [SIZE][SIZE]int
	dx     = []int{-1, 1, 0, 0}
	dy     = []int{0, 0, -1, 1}
)

func Day12(input []string) {
	for r := 0; r < len(input); r++ {
		for c := 0; c < len(input[r]); c++ {
			grid[r][c] = input[r][c]
		}
	}

	shapeID := 1

	for r := 0; r < SIZE; r++ {
		for c := 0; c < SIZE; c++ {
			if shapes[r][c] == 0 {
				floodFill(r, c, shapeID, grid[r][c])
				shapeID++
			}
		}
	}

	p1Price := 0
	p2Price := 0

	for id := 1; id < shapeID; id++ {
		p1Price += getPrice(id)
		p2Price += part2(id)
	}

	fmt.Println("Part 1")
	fmt.Println(p1Price)
	fmt.Println("Part 2")
	fmt.Println(p2Price)
}

func floodFill(x, y, id int, plant byte) {
	stack := []Point{{x, y}}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if shapes[curr.x][curr.y] != 0 {
			continue
		}

		shapes[curr.x][curr.y] = id

		for i := 0; i < 4; i++ {
			nx, ny := curr.x+dx[i], curr.y+dy[i]

			if nx >= 0 && ny >= 0 && nx < SIZE && ny < SIZE && shapes[nx][ny] == 0 && grid[nx][ny] == plant {
				stack = append(stack, Point{nx, ny})
			}
		}
	}
}

func getPrice(id int) int {
	area, perimeter := 0, 0

	for r := range SIZE {
		inShape := false

		for c := range SIZE {
			if shapes[r][c] == id {
				area++

				if !inShape {
					inShape = true
					perimeter++
				}
			} else if inShape {
				inShape = false
				perimeter++
			}
		}

		if inShape {
			perimeter++
		}
	}

	for c := range SIZE {
		inShape := false

		for r := range SIZE {
			if shapes[r][c] == id {
				if !inShape {
					inShape = true
					perimeter++
				}
			} else if inShape {
				inShape = false
				perimeter++
			}
		}

		if inShape {
			perimeter++
		}
	}

	return area * perimeter
}

func part2(id int) int {
	var sides, area int

	hChanges := make([][]int, SIZE)
	vChanges := make([][]int, SIZE+1)

	for i := range hChanges {
		hChanges[i] = make([]int, SIZE+1)
	}

	for i := range vChanges {
		vChanges[i] = make([]int, SIZE)
	}

	for r := 0; r < SIZE; r++ {
		inShape := 0

		for c := 0; c < SIZE; c++ {
			if shapes[r][c] == id {
				area++

				if inShape == 0 {
					inShape = 1
					hChanges[r][c] = 1

					if r == 0 || hChanges[r-1][c] != 1 {
						sides++
					}
				}
			} else if inShape == 1 {
				inShape = 0
				hChanges[r][c] = 2
				if r == 0 || hChanges[r-1][c] != 2 {
					sides++
				}
			}
		}

		if inShape == 1 {
			hChanges[r][SIZE] = 2
			if r == 0 || hChanges[r-1][SIZE] != 2 {
				sides++
			}
		}
	}

	for c := 0; c < SIZE; c++ {
		inShape := 0
		for r := 0; r < SIZE; r++ {
			if shapes[r][c] == id {
				if inShape == 0 {
					inShape = 1
					vChanges[r][c] = 1

					if c == 0 || vChanges[r][c-1] != 1 {
						sides++
					}
				}
			} else if inShape == 1 {
				inShape = 0
				vChanges[r][c] = 2

				if c == 0 || vChanges[r][c-1] != 2 {
					sides++
				}
			}
		}

		if inShape == 1 {
			vChanges[SIZE][c] = 2
			if c == 0 || vChanges[SIZE][c-1] != 2 {
				sides++
			}
		}
	}

	return area * sides
}
