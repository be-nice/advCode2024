package day11

import (
	"fmt"
	"strconv"
)

func Day11(n []int) {
	stones := make(map[int]int)
	part1 := 0
	part2 := 0

	for _, num := range n {
		stones[num]++
	}

	for i := range 75 {
		if i == 25 {
			for _, val := range stones {
				part1 += val
			}
		}

		newStones := make(map[int]int)

		for stone, val := range stones {
			if stone == 0 {
				newStones[1] += val
			} else if len(strconv.Itoa(stone))%2 == 0 {
				full := strconv.Itoa(stone)
				mid := len(full) / 2
				left, _ := strconv.Atoi(full[:mid])
				right, _ := strconv.Atoi(full[mid:])
				newStones[left] += val
				newStones[right] += val
			} else {
				newStones[stone*2024] += val
			}
		}

		stones = newStones
	}

	for _, val := range stones {
		part2 += val
	}

	fmt.Println("Part 1")
	fmt.Println(part1)
	fmt.Println("Part 2")
	fmt.Println(part2)
}
