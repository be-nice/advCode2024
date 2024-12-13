package day10

import "fmt"

type point struct {
	x, y int
}

var directions = []point{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
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

	for _, dir := range directions {
		newX := curr.x + dir.x
		newY := curr.y + dir.y

		if newX >= 0 && newX < len(input[0]) && newY >= 0 && newY < len(input) {
			if input[newY][newX] == input[curr.y][curr.x]+1 {
				validStep = append(validStep, point{newX, newY})
			}
		}
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
