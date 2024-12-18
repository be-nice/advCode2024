package main

import (
	"adv/day1"
	"adv/day10"
	"adv/day11"
	"adv/day12"
	"adv/day13"
	"adv/day14"
	"adv/day15"
	"adv/day16"
	"adv/day17"
	"adv/day18"
	"adv/day2"
	"adv/day3"
	"adv/day4"
	"adv/day5"
	"adv/day6"
	"adv/day7"
	"adv/day8"
	"adv/day9"
	"adv/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type day struct {
	name string
	fn   func()
}

var solutions []day

func init() {
	solutions = []day{
		{"Day 1 solution:", func() {
			day1.Day1(ReadInputLnStr("./day1/input.txt"))
		}},
		{"Day 2 solution:", func() {
			day2.Day2(ReadInputLnStr("./day2/input.txt"))
		}},
		{"Day 3 solution:", func() {
			day3.Day3(ReadInputBlockByte("./day3/input.txt"))
		}},
		{"Day 4 solution:", func() {
			day4.Day4(ReadInputLnStr("./day4/input.txt"))
		}},
		{"Day 5 solution:", func() {
			day5.Day5(ReadInputBlockByte("./day5/input.txt"))
		}},
		{"Day 6 solution:", func() {
			day6.Day6(ReadInputLineRune("./day6/input.txt"))
		}},
		{"Day 7 solution:", func() {
			day7.Day7(ReadEquation("./day7/input.txt"))
		}},
		{"Day 8 solution:", func() {
			day8.Day8(ReadInputLnStr("./day8/input.txt"))
		}},
		{"Day 9 solution:", func() {
			day9.Day9(ReadInputBlockString("./day9/input.txt"))
		}},
		{"Day 10 solution:", func() {
			day10.Day10(ReadInputLnInt("./day10/input.txt"))
		}},
		{"Day 11 solution:", func() {
			day11.Day11(ReadInputBlockInt("./day11/input.txt"))
		}},
		{"Day 12 solution:", func() {
			day12.Day12(ReadInputLnStr("./day12/input.txt"))
		}},
		{"Day 13 solution:", func() {
			day13.Day13(ReadInputLnStr("./day13/input.txt"))
		}},
		{"Day 14 solution:", func() {
			day14.Day14(ReadInputLnStr("./day14/input.txt"))
		}},
		{"Day 15 solution:", func() {
			day15.Day15(ReadInputBlockString("./day15/input.txt"))
		}},
		{"Day 16 solution:", func() {
			day16.Day16(ReadInputLineRune("./day16/input.txt"))
		}},
		{"Day 17 solution:", func() {
			day17.Day17(ReadInputLnStr("./day17/input.txt"))
		}},
		{"Day 18 solution:", func() {
			day18.Day18(ReadInputLnStr("./day18/input.txt"))
		}},
	}
}

func main() {
	t := &utils.Timer{}

	switch len(os.Args) {
	case 1:
		for _, day := range solutions {
			fmt.Println(day.name)

			t.StartTimer()
			day.fn()
			t.PrintDuration()

			fmt.Println("----------------------------------------------------")
		}
		t.PrintTotalDuration()
	case 2:
		idx, err := strconv.Atoi(os.Args[1])
		if err != nil || idx < 1 || idx > len(solutions) {
			fmt.Printf("Expected [<int> << range(1, %d)] || Got %s\n", len(solutions), os.Args[1])
			os.Exit(1)
		}

		fmt.Println(solutions[idx-1].name)

		t.StartTimer()
		solutions[idx-1].fn()
		t.PrintDuration()
	default:
		fmt.Printf("Expected arguments 0 or 1 [<int> <- range(1, %d)] || Got %d arguments\n", len(solutions), len(os.Args)-1)
		os.Exit(1)
	}
}

func ReadInputLnStr(dir string) []string {
	res := []string{}

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		res = append(res, scanner.Text())
	}

	return res
}

func ReadInputLnByte(dir string) [][]byte {
	res := [][]byte{}

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		res = append(res, scanner.Bytes())
	}

	return res
}

func ReadInputLnInt(dir string) [][]int {
	res := [][]int{}

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		in := scanner.Bytes()
		nums := make([]int, len(in))
		for i, val := range in {
			nums[i] = int(val - '0')
		}
		res = append(res, nums)
	}

	return res
}

func ReadInputBlockInt(dir string) []int {
	var res []int

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	dataStr := scanner.Text()
	splitData := strings.Fields(dataStr)

	for _, val := range splitData {
		num, _ := strconv.Atoi(val)
		res = append(res, num)
	}

	return res
}

func ReadInputBlockByte(dir string) []byte {
	data, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
}

func ReadInputBlockString(dir string) string {
	data, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func ReadInputLineRune(dir string) [][]rune {
	res := [][]rune{}

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		res = append(res, []rune(scanner.Text()))
	}

	return res
}

func ReadEquation(dir string) []day7.Equation {
	var res []day7.Equation

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		result, _ := strconv.Atoi(lineSplit[0])
		var listOfNums []int
		listOfNumsAsSting := strings.Split(strings.TrimSpace(lineSplit[1]), " ")

		for i := range listOfNumsAsSting {
			num, _ := strconv.Atoi(listOfNumsAsSting[i])
			listOfNums = append(listOfNums, num)
		}

		res = append(res, day7.Equation{Res: result, Calibration: listOfNums})
	}

	return res
}
