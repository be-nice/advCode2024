package day11

import (
	"fmt"
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
			} else {
				digits := 0
				temp := stone

				for temp > 0 {
					temp /= 10
					digits++
				}

				if digits%2 == 0 {
					halfDigits := digits / 2
					divisor := 1

					for range halfDigits {
						divisor *= 10
					}

					left := stone / divisor
					right := stone % divisor
					newStones[left] += val
					newStones[right] += val
				} else {
					newStones[stone*2024] += val
				}
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
