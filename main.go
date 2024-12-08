package main

import (
	"adv/day1"
	"adv/day2"
	"adv/day3"
	"adv/day4"
	"adv/day5"
	"adv/day6"
	"adv/day8"
	"adv/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		{"Day 8 solution:", func() {
			day8.Day8(ReadInputLnStr("./day8/input.txt"))
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

func ReadInputBlockByte(dir string) []byte {
	data, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data
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
