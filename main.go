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
)

func main() {
	t := &utils.Timer{}

	fmt.Println("Day 1 solution:")
	t.StartTimer()
	day1.Day1(ReadInputLnStr("./day1/input.txt"))
	t.PrintDuration()
	fmt.Println("----------------------------------------------------")

	fmt.Println("Day 2 solution:")
	t.StartTimer()
	day2.Day2(ReadInputLnStr("./day2/input.txt"))
	t.PrintDuration()

	fmt.Println("----------------------------------------------------")

	fmt.Println("Day 3 solution:")
	t.StartTimer()
	day3.Day3(ReadInputBlockByte("./day3/input.txt"))
	t.PrintDuration()

	fmt.Println("-----------------------------------------------------")

	fmt.Println("Day 4 solution:")
	t.StartTimer()
	day4.Day4(ReadInputLnStr("./day4/input.txt"))
	t.PrintDuration()

	fmt.Println("-----------------------------------------------------")

	fmt.Println("Day 5 solution:")
	t.StartTimer()
	day5.Day5(ReadInputBlockByte("./day5/input.txt"))
	t.PrintDuration()

	fmt.Println("-----------------------------------------------------")

	fmt.Println("Day 6 solution:")
	t.StartTimer()
	day6.Day6(ReadInputLineRune("./day6/input.txt"))
	t.PrintDuration()

	fmt.Println("-----------------------------------------------------")

	fmt.Println("Day 8 solution;")
	t.StartTimer()
	day8.Day8(ReadInputLnStr("./day8/input.txt"))
	t.PrintDuration()
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
