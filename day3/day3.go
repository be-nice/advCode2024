package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

func Day3(b []byte) {
	sum := 0

	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllSubmatch(b, -1)

	for _, match := range matches {
		n1, _ := strconv.Atoi(string(match[1]))
		n2, _ := strconv.Atoi(string(match[2]))
		sum += n1 * n2
	}

	fmt.Println("Part 1")
	fmt.Println(sum)

	part2(b)
}

func part2(b []byte) {
	sum := 0

	re := regexp.MustCompile(`(mul\((\d+),(\d+)\)|do\(\)|don't\(\))`)
	matches := re.FindAllSubmatch(b, -1)

	allowed := true
	for _, match := range matches {
		if string(match[0]) == "do()" {
			allowed = true
			continue
		}

		if string(match[0]) == "don't()" {
			allowed = false
			continue
		}

		if allowed {
			n1, _ := strconv.Atoi(string(match[2]))
			n2, _ := strconv.Atoi(string(match[3]))
			sum += n1 * n2
		}
	}

	fmt.Println("Part 2")
	fmt.Println(sum)
}
