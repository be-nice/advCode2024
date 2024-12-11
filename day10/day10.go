package day10

import "fmt"

type point struct {
	x, y int
}

func Day10(n [][]int) {
	part1, part2 := findHeads(n)

	fmt.Println("Part 1")
	fmt.Println(part1)
	fmt.Println("Part 2")
	fmt.Println(part2)
}

func findNext(input [][]int, curr point) []point {
	validStep := []point{}

	if curr.x > 0 && input[curr.y][curr.x-1] == input[curr.y][curr.x]+1 {
		validStep = append(validStep, point{curr.x - 1, curr.y})
	}

	if curr.x < len(input[0])-1 && input[curr.y][curr.x+1] == input[curr.y][curr.x]+1 {
		validStep = append(validStep, point{curr.x + 1, curr.y})
	}

	if curr.y > 0 && input[curr.y-1][curr.x] == input[curr.y][curr.x]+1 {
		validStep = append(validStep, point{curr.x, curr.y - 1})
	}

	if curr.y < len(input)-1 && input[curr.y+1][curr.x] == input[curr.y][curr.x]+1 {
		validStep = append(validStep, point{curr.x, curr.y + 1})
	}

	return validStep
}

func findScore(input [][]int, start point, trailHeads map[point]struct{}, count int) (map[point]struct{}, int) {
	if input[start.y][start.x] == 9 {
		if _, ok := trailHeads[start]; !ok {
			trailHeads[start] = struct{}{}
		}

		return trailHeads, count + 1
	}

	nextSteps := findNext(input, start)

	if len(nextSteps) == 0 {
		return trailHeads, count
	}

	for _, step := range nextSteps {
		trailHeads, count = findScore(input, step, trailHeads, count)
	}

	return trailHeads, count
}

func findHeads(input [][]int) (int, int) {
	sumScore := 0
	sumRating := 0

	for j, row := range input {
		for i, char := range row {
			if char == 0 {
				score, rating := findScore(input, point{i, j}, make(map[point]struct{}), 0)
				sumScore += len(score)
				sumRating += rating
			}
		}
	}

	return sumScore, sumRating
}
