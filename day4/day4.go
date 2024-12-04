package day4

import (
	"fmt"
)

func Day4(s []string) {
	count := 0
	masCount := 0

	for i, row := range s {
		for j, ch := range row {
			if ch == 'X' {
				count += checkMatch(s, i, j)
			}
			if i+1 < len(s)-1 && s[i+1][j] == byte('A') && j < len(row) {
				masCount += part2(s, i+1, j)
			}
		}
	}

	fmt.Println("Part 1")
	fmt.Println(count)
	fmt.Println("Part 2")
	fmt.Println(masCount)
}

func checkMatch(s []string, i, j int) int {
	count := 0

	if j >= 3 {
		if s[i][j-1] == 'M' && s[i][j-2] == 'A' && s[i][j-3] == 'S' {
			count++
		}
	}

	if j <= (len(s[0]) - 1 - 3) {
		if s[i][j+1] == 'M' && s[i][j+2] == 'A' && s[i][j+3] == 'S' {
			count++
		}
	}

	if i >= 3 {
		if s[i-1][j] == 'M' && s[i-2][j] == 'A' && s[i-3][j] == 'S' {
			count++
		}
	}

	if i <= (len(s[0]) - 1 - 3) {
		if s[i+1][j] == 'M' && s[i+2][j] == 'A' && s[i+3][j] == 'S' {
			count++
		}
	}

	if i >= 3 && j <= (len(s[0])-1-3) {
		if s[i-1][j+1] == 'M' && s[i-2][j+2] == 'A' && s[i-3][j+3] == 'S' {
			count++
		}
	}

	if i >= 3 && j >= 3 {
		if s[i-1][j-1] == 'M' && s[i-2][j-2] == 'A' && s[i-3][j-3] == 'S' {
			count++
		}
	}

	if i <= (len(s[0])-1-3) && j <= (len(s[0])-1-3) {
		if s[i+1][j+1] == 'M' && s[i+2][j+2] == 'A' && s[i+3][j+3] == 'S' {
			count++
		}
	}

	if i <= (len(s[0])-1-3) && j >= 3 {
		if s[i+1][j-1] == 'M' && s[i+2][j-2] == 'A' && s[i+3][j-3] == 'S' {
			count++
		}
	}

	return count
}

func part2(s []string, i, j int) int {
	count := 0

	if j < 1 || j > len(s[i])-2 {
		return 0
	}

	if s[i-1][j+1] == 'M' && s[i-1][j-1] == 'M' && s[i+1][j-1] == 'S' && s[i+1][j+1] == 'S' {
		count++
	}

	if s[i-1][j+1] == 'S' && s[i-1][j-1] == 'S' && s[i+1][j-1] == 'M' && s[i+1][j+1] == 'M' {
		count++
	}

	if s[i-1][j+1] == 'S' && s[i-1][j-1] == 'M' && s[i+1][j-1] == 'M' && s[i+1][j+1] == 'S' {
		count++
	}

	if s[i-1][j+1] == 'M' && s[i-1][j-1] == 'S' && s[i+1][j-1] == 'S' && s[i+1][j+1] == 'M' {
		count++
	}
	return count
}

// Fun solution, but perf sucks hard

//func Day4(s []string) {
//	count := 0
//	masCount := 0
//
//	grid := map[image.Point]rune{}
//	for y, row := range s {
//		for x, ch := range row {
//			grid[image.Point{x, y}] = ch
//		}
//	}
//
//	adj := func(p image.Point, pd int) []string {
//		direction := []image.Point{
//			{-1, -1},
//			{1, -1},
//			{1, 1},
//			{-1, 1},
//			{0, -1},
//			{1, 0},
//			{0, 1},
//			{-1, 0},
//		}
//
//		words := make([]string, len(direction))
//		for i, d := range direction {
//			for n := range pd {
//				words[i] += string(grid[p.Add(d.Mul(n))])
//			}
//		}
//		return words
//	}
//
//	for p := range grid {
//		count += strings.Count(strings.Join(adj(p, 4), " "), "XMAS")
//		masCount += strings.Count("AMAMASASAMAMAS", strings.Join(adj(p, 2)[:4], ""))
//	}
//	fmt.Println("Part 1")
//	fmt.Println(count)
//	fmt.Println("Part 2")
//	fmt.Println(masCount)
//}
