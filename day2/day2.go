package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2(s []string) {
	safeCount := 0
	advSafeCount := 0

	nums := parseInt(s)

	for _, row := range nums {
		if validSeq(row) {
			safeCount++
		} else {
			for i := range len(row) {
				permT := make([]int, i, len(row)-1)
				copy(permT, row[:i])
				permT = append(permT, row[i+1:]...)
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

func parseInt(s []string) [][]int {
	res := make([][]int, 0, len(s))

	for _, line := range s {
		nums := strings.Fields(line)
		numArr := make([]int, 0, len(nums))

		for _, val := range nums {
			num, _ := strconv.Atoi(val)
			numArr = append(numArr, num)
		}

		res = append(res, numArr)
	}

	return res
}

func validSeq(nums []int) bool {
	l1 := nums[0]
	l2 := nums[1]
	dir := l1 < l2
	for i := range len(nums) - 1 {
		l1 := nums[i]
		l2 := nums[i+1]

		validDiff := l1-l2 >= -3 && l1-l2 <= 3
		if l1 == l2 || l1 < l2 != dir || !validDiff {
			return false
		}
	}
	return true
}

//func validSeq(nums []string) bool {
//	diffMap := map[int]int{1: 0, 2: 0, 3: 0}
//	negDiffMap := map[int]int{-1: 0, -2: 0, -3: 0}
//	for i := range len(nums) - 1 {
//		n, _ := strconv.Atoi(nums[i])
//		j, _ := strconv.Atoi(nums[i+1])
//		diffMap[n-j]++
//		negDiffMap[n-j]++
//	}
//
//	return len(diffMap) == 3 || len(negDiffMap) == 3
//}
