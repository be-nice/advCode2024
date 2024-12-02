package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2(s []string) {
	safeCount := 0
	advSafeCount := 0

	for _, row := range s {
		nums := strings.Fields(row)
		if validSeq(nums) {
			safeCount++
		} else {
			for i := range len(nums) {
				permT := make([]string, i, len(nums)-1)
				copy(permT, nums[:i])
				permT = append(permT, nums[i+1:]...)
				if validSeq(permT) {
					advSafeCount++
					break
				}
			}
		}
	}

	fmt.Println("Part 1")
	fmt.Println(safeCount)
	fmt.Println("Part 2")
	fmt.Println(safeCount + advSafeCount)
}

func validSeq(nums []string) bool {
	diffMap := map[int]int{1: 0, 2: 0, 3: 0}
	negDiffMap := map[int]int{-1: 0, -2: 0, -3: 0}
	for i := range len(nums) - 1 {
		n, _ := strconv.Atoi(nums[i])
		j, _ := strconv.Atoi(nums[i+1])
		diffMap[n-j]++
		negDiffMap[n-j]++
	}

	return len(diffMap) == 3 || len(negDiffMap) == 3
}
