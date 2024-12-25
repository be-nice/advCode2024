package day25

import (
	"fmt"
	"strings"
	"sync"
)

const TOTALHEIGHT = 7

var mu sync.Mutex

func Day25(s string) {
	keys, locks := parse(s)
	lockHeights := make([][]int, 0, len(locks))
	keyHeights := make([][]int, 0, len(keys))
	keyChan := make(chan []int, 100)
	lockChan := make(chan []int, 100)
	countChan := make(chan struct{}, 5000)
	var wg sync.WaitGroup
	var wwg sync.WaitGroup
	var lwg sync.WaitGroup

	wwg.Add(1)
	go func() {
		defer wwg.Done()
		for val := range keyChan {
			keyHeights = append(keyHeights, val)
		}
	}()

	wwg.Add(1)
	go func() {
		defer wwg.Done()
		for val := range lockChan {
			lockHeights = append(lockHeights, val)
		}
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
		defer wg.Done()
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
	wwg.Wait()

	for _, lock := range lockHeights {
		lwg.Add(1)
		go func(lock []int) {
			defer lwg.Done()
			for _, key := range keyHeights {
				checkFit(lock, key, countChan)
			}
		}(lock)
	}

	lwg.Wait()
	close(countChan)

	fmt.Println("Part 1")
	fmt.Println(len(countChan))
}

func parse(s string) ([][]string, [][]string) {
	keys := make([][]string, 0, len(s)/2)
	locks := make([][]string, 0, len(s)/2)

	for _, block := range strings.Split(s, "\n\n") {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		lock := lines[0] == "#####"

		if lock {
			locks = append(locks, lines)
		} else {
			keys = append(keys, lines)
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

func checkFit(lock, key []int, countChan chan struct{}) {
	for i := range len(lock) {
		if lock[i]+key[i] > TOTALHEIGHT {
			return
		}
	}

	countChan <- struct{}{}
}
