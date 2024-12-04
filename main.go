package main

import (
	"adv/day1"
	"adv/day2"
	"adv/day3"
	"adv/day4"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Day 1 solution:")
	start := time.Now()
	day1.Day1(ReadInputLnStr("./day1/input.txt"))
	fmt.Println(time.Since(start).Seconds())

	fmt.Println("----------------------------------------------------")

	fmt.Println("Day 2 solution:")
	start = time.Now()
	day2.Day2(ReadInputLnStr("./day2/input.txt"))
	fmt.Println(time.Since(start).Seconds())

	fmt.Println("----------------------------------------------------")

	fmt.Println("Day 3 solution:")
	start = time.Now()
	day3.Day3(ReadInputBlockByte("./day3/input.txt"))
	fmt.Println(time.Since(start).Seconds())

	fmt.Println("-----------------------------------------------------")

	fmt.Println("Day 4 solution:")
	start = time.Now()
	day4.Day4(ReadInputLnStr("./day4/input.txt"))
	fmt.Println(time.Since(start).Seconds())
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
