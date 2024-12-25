package day25

import (
	"fmt"
	"strings"
	"sync"
)

func Day25(s string) {
	keys, locks := parse(s)
	totalHeight := 7
	count := 0
	lockHeights := make([][]int, 0, len(locks))
	keyHeights := make([][]int, 0, len(keys))
	keyChan := make(chan []int, 100)
	lockChan := make(chan []int, 100)
	syncChan := make(chan struct{}, 2)
	var wg sync.WaitGroup

	go func() {
		for val := range keyChan {
			keyHeights = append(keyHeights, val)
		}
		syncChan <- struct{}{}
	}()

	go func() {
		for val := range lockChan {
			lockHeights = append(lockHeights, val)
		}
		syncChan <- struct{}{}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, key := range keys {
			wg.Add(1)
			go func(key []string) {
				defer wg.Done()
				getHeights(key, false, keyChan)
			}(key)
		}
	}()
	wg.Add(1)
	go func() {
		wg.Done()
		for _, lock := range locks {
			wg.Add(1)
			go func(lock []string) {
				defer wg.Done()
				getHeights(lock, true, lockChan)
			}(lock)
		}
	}()

	wg.Wait()
	close(keyChan)
	close(lockChan)

	sync := 0
	for range syncChan {
		sync++
		if sync == 2 {
			break
		}
	}

	for _, lock := range lockHeights {
		for _, key := range keyHeights {
			if checkFit(lock, key, totalHeight) {
				count++
			}
		}
	}

	fmt.Println("Part 1")
	fmt.Println(count)
}

func parse(s string) ([][]string, [][]string) {
	keys := make([][]string, 0, len(s)/2)
	locks := make([][]string, 0, len(s)/2)

	for _, block := range strings.Split(s, "\n\n") {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		lock := lines[0] == "#####"

		t := make([]string, 0, len(lines))

		for _, line := range lines {
			t = append(t, line)
		}

		if lock {
			locks = append(locks, t)
		} else {
			keys = append(keys, t)
		}
	}

	return keys, locks
}

func getHeights(schematic []string, isLock bool, send chan []int) {
	heights := make([]int, len(schematic[0]))

	for i := range len(schematic[0]) {
		height := 0

		if isLock {
			for j := range len(schematic) {
				if schematic[j][i] == '#' {
					height++
				} else {
					break
				}
			}
		} else {
			for j := len(schematic) - 1; j >= 0; j-- {
				if schematic[j][i] == '#' {
					height++
				} else {
					break
				}
			}
		}

		heights[i] = height
	}

	send <- heights
}

func checkFit(lock, key []int, totalHeight int) bool {
	for i := range len(lock) {
		if lock[i]+key[i] > totalHeight {
			return false
		}
	}

	return true
}
