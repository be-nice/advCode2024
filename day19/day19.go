package day19

import (
	"fmt"
	"strings"
	"sync"
)

var (
	cache   = make(map[string]int)
	cacheMu sync.Mutex
)

func Day19(s string) {
	patterns, partials := parse(s)
	res := 0
	perms := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, line := range patterns {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			temp := dp(partials, line)

			mu.Lock()
			if temp > 0 {
				res++
			}
			perms += temp
			mu.Unlock()
		}(line)
	}

	wg.Wait()

	fmt.Println("Part 1")
	fmt.Println(res)
	fmt.Println("Part 2")
	fmt.Println(perms)
}

func dp(partials []string, line string) int {
	cacheMu.Lock()
	if result, ok := cache[line]; ok {
		cacheMu.Unlock()
		return result
	}
	cacheMu.Unlock()

	if len(line) == 0 {
		return 1
	}

	res := 0
	for _, k := range partials {
		if strings.HasPrefix(line, k) {
			res += dp(partials, line[len(k):])
		}
	}

	cacheMu.Lock()
	cache[line] = res
	cacheMu.Unlock()

	return res
}

func parse(s string) ([]string, []string) {
	split := strings.Split(s, "\n\n")
	lines := strings.Split(strings.TrimSpace(split[1]), "\n")
	return lines, strings.Split(split[0], ", ")
}
