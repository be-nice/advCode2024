package day19

import (
	"fmt"
	"strings"
)

var cache = make(map[string]int)

func Day19(s string) {
	patterns, partials := parse(s)
	res := 0
	perms := 0

	for _, line := range patterns {
		temp := dp(partials, line)

		if temp > 0 {
			res++
		}

		perms += temp
	}

	fmt.Println("Part 1")
	fmt.Println(res)
	fmt.Println("Part 2")
	fmt.Println(perms)
}

func dp(partials []string, line string) int {
	if _, ok := cache[line]; !ok {
		if len(line) == 0 {
			return 1
		} else {
			res := 0

			for _, k := range partials {
				if strings.HasPrefix(line, k) {
					res += dp(partials, line[len(k):])
				}
			}

			cache[line] = res
		}
	}

	return cache[line]
}

func parse(s string) ([]string, []string) {
	split := strings.Split(s, "\n\n")

	lines := strings.Split(strings.TrimSpace(split[1]), "\n")

	return lines, strings.Split(split[0], ", ")
}
