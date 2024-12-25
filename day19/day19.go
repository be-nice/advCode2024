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
	resChan := make(chan int, len(patterns))
	var wg sync.WaitGroup
	var wwg sync.WaitGroup

	wwg.Add(1)
	go func() {
		defer wwg.Done()
		for val := range resChan {
			if val > 0 {
				res++
			}
			perms += val
		}
	}()

	for _, line := range patterns {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			resChan <- dp(partials, line)
		}(line)
	}

	wg.Wait()
	close(resChan)
	wwg.Wait()

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
