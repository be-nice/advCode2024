package day5

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day5(b []byte) {
	split := strings.Split(strings.TrimSpace(string(b)), "\n\n")

	eval := func(sorted bool) int {
		sum := 0
		for _, line := range strings.Split(split[1], "\n") {
			if ruleOrd := strings.Split(line, ","); slices.IsSortedFunc(ruleOrd, func(a, b string) int {
				return compare(a, b, split[0])
			}) == sorted {
				slices.SortFunc(ruleOrd, func(a, b string) int {
					return compare(a, b, split[0])
				})
				n, _ := strconv.Atoi(ruleOrd[len(ruleOrd)/2])
				sum += n
			}
		}
		return sum
	}

	fmt.Println("Part 1")
	fmt.Println(eval(true))

	fmt.Println("Part 2")
	fmt.Println(eval(false))
}

func compare(a, b string, rulesList string) int {
	for _, row := range strings.Split(rulesList, "\n") {
		if ruleOrd := strings.Split(row, "|"); ruleOrd[0] == a && ruleOrd[1] == b {
			return -1
		}
	}
	return 0
}
