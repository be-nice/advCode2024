package day1

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func Day1(s []string) {
	left := make([]int, 0, len(s))
	right := make([]int, 0, len(s))

	for _, row := range s {
		l, _ := strconv.Atoi(strings.Fields(row)[0])
		r, _ := strconv.Atoi(strings.Fields(row)[1])
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i, val := range left {
		sum += int(math.Abs(float64(val - right[i])))
	}

	fmt.Println("Part 1")
	fmt.Println(sum)

	part2(left, right)
}

func part2(left, right []int) {
	freqMap := make(map[int]int, len(right))

	for _, val := range right {
		freqMap[val]++
	}

	sum := 0

	for _, val := range left {
		sum += freqMap[val] * val
	}

	fmt.Println("part2")
	fmt.Println(sum)
}
