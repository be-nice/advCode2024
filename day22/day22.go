package day22

import (
	"fmt"
	"math"
	"sync"
)

const MASK = (1 << 24) - 1

type Result struct {
	finalVal int
	steps    []int
}

func Day22(n []int) {
	var wg sync.WaitGroup
	resChan := make(chan Result, len(n))
	sum := 0
	totalSteps := make([][]int, 0, len(n))
	syncChan := make(chan struct{})

	go func() {
		for val := range resChan {
			sum += val.finalVal
			totalSteps = append(totalSteps, val.steps)
		}

		syncChan <- struct{}{}
	}()

	for _, val := range n {
		wg.Add(1)
		go func(n int, c chan Result) {
			defer wg.Done()
			part1(n, c)
		}(val, resChan)
	}

	wg.Wait()
	close(resChan)
	<-syncChan

	fmt.Println("Part 1")
	fmt.Println(sum)
	fmt.Println("Part 2")
	fmt.Println(part2(totalSteps))
}

func part1(n int, c chan Result) {
	steps := make([]int, 2001)

	for i := range 2001 {
		n = (n ^ n<<6) & MASK
		n ^= n >> 5
		n = (n ^ n<<11) & MASK

		steps[i] = n
	}
	c <- Result{finalVal: steps[2000-1], steps: steps}
}

func idx(a, b, c, d int) int {
	return 6859*(a+9) + 361*(b+9) + 19*(c+9) + (d + 9)
}

func part2(allSteps [][]int) int {
	size := int(math.Pow(19, 4))
	dp := make([]int, size)
	dpi := make([]int, size)
	maxValue := 0

	for i, steps := range allSteps {
		for j := 4; j < 2001; j++ {
			e, d, c, b, a := steps[j-4]%10, steps[j-3]%10, steps[j-2]%10, steps[j-1]%10, steps[j]%10
			index := idx(d-e, c-d, b-c, a-b)
			if dpi[index] != i+1 {
				dpi[index] = i + 1
				dp[index] += a
			}
		}
	}

	for _, v := range dp {
		if v > maxValue {
			maxValue = v
		}
	}

	return maxValue
}
