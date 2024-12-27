package day13

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type clawMachine struct {
	btnA, btnB struct {
		x, y int
	}
	prize struct {
		c, d int
	}
}

func Day13(s string) {
	data := make([]clawMachine, 0, len(s)>>3)
	var wg sync.WaitGroup
	var wwg sync.WaitGroup
	resChan := make(chan clawMachine, len(s)>>3)

	wwg.Add(1)
	go func() {
		defer wwg.Done()
		for val := range resChan {
			data = append(data, val)
		}
	}()

	for _, block := range strings.Split(strings.TrimSpace(s), "\n\n") {
		wg.Add(1)
		go func(block string) {
			defer wg.Done()

			parts := strings.Split(strings.TrimSpace(block), "\n")
			aParts := strings.Split(parts[0], " ")
			bParts := strings.Split(parts[1], " ")

			ax, _ := strconv.Atoi(aParts[2][2 : len(aParts[2])-1])
			ay, _ := strconv.Atoi(aParts[3][2:])
			bx, _ := strconv.Atoi(bParts[2][2 : len(bParts[2])-1])
			by, _ := strconv.Atoi(bParts[3][2:])

			tParts := strings.Split(parts[2], " ")
			c, _ := strconv.Atoi(tParts[1][2 : len(tParts[1])-1])
			d, _ := strconv.Atoi(tParts[2][2:])

			curr := clawMachine{
				btnA:  struct{ x, y int }{x: ax, y: ay},
				btnB:  struct{ x, y int }{x: bx, y: by},
				prize: struct{ c, d int }{c: c, d: d},
			}

			resChan <- curr
		}(block)
	}

	wg.Wait()
	close(resChan)
	wwg.Wait()

	fmt.Println("Part 1")
	fmt.Println(solve(data, false))
	fmt.Println("Part 2")
	fmt.Println(solve(data, true))
}

func solve(data []clawMachine, part2 bool) int {
	tokens := 0
	add := 0

	if part2 {
		add = 10000000000000
	}

	for _, val := range data {
		ax, ay := val.btnA.x, val.btnA.y
		bx, by := val.btnB.x, val.btnB.y
		tx := val.prize.c + add
		ty := val.prize.d + add

		a := float64(tx*by-ty*bx) / float64(ax*by-ay*bx)
		b := float64(ty*ax-tx*ay) / float64(ax*by-ay*bx)

		if a == float64(int(a)) && b == float64(int(b)) {
			tokens += 3*int(a) + int(b)
		}
	}

	return tokens
}
