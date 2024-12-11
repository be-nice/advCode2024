package day7

import (
	"fmt"
	"math"
)

type Equation struct {
	Res         int
	Calibration []int
}

func Day7(equations []Equation) {
	total := 0
	part2 := 0

	for _, val := range equations {
		if isValid(val, false) {
			total += val.Res
		} else if isValid(val, true) {
			part2 += val.Res
		}
	}

	fmt.Println("Part 1")
	fmt.Println(total)
	fmt.Println("Part 2")
	fmt.Println(total + part2)
}

func isValid(e Equation, concat bool) bool {
	concatRes := false

	addRes := tryCombR(e.Calibration, 0, e.Res, "+", concat)
	mulRes := tryCombR(e.Calibration, 0, e.Res, "*", concat)
	if concat {
		concatRes = tryCombR(e.Calibration, 0, e.Res, "||", concat)
	}

	return addRes || mulRes || concatRes
}

func tryCombR(nums []int, total int, target int, operation string, concat bool) bool {
	newTotal := 0

	switch operation {
	case "+":
		newTotal = total + nums[0]
	case "*":
		newTotal = total * nums[0]
	case "||":
		digitCount := 0
		temp := nums[0]
		for temp > 0 {
			digitCount++
			temp /= 10
		}
		newTotal = total*int(math.Pow(10, float64(digitCount))) + nums[0]
	}

	if len(nums) == 1 || newTotal > target {
		return newTotal == target
	}

	concatRes := false

	addRes := tryCombR(nums[1:], newTotal, target, "+", concat)
	mulRes := tryCombR(nums[1:], newTotal, target, "*", concat)

	if concat {
		concatRes = tryCombR(nums[1:], newTotal, target, "||", concat)
	}

	return addRes || mulRes || concatRes
}
