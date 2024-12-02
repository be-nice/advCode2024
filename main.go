package main

import (
	"adv/day1"
	"adv/day2"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// fmt.Println("Day 1 test input:")
	// day1.Day1(ReadTestInputStr("day1"))

	fmt.Println("Day 1 solution:")
	start := time.Now()
	day1.Day1(ReadMainInputStr("day1"))
	fmt.Println(time.Since(start))

	fmt.Println("----------------------------------------------------")

	// fmt.Println("Day 2 test input:")
	// day2.Day2(ReadTestInputStr("day2"))

	fmt.Println("Day 2 solution:")
	start = time.Now()
	day2.Day2(ReadMainInputStr("day2"))
	fmt.Println(time.Since(start))
}

func ReadTestInputStr(dir string) []string {
	res := []string{}

	file, err := os.Open(filepath.Join(dir, "test.txt"))
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

func ReadMainInputStr(dir string) []string {
	res := []string{}

	file, err := os.Open(filepath.Join(dir, "main.txt"))
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

func ReadTestInputByte(dir string) [][]byte {
	res := [][]byte{}

	file, err := os.Open(filepath.Join(dir, "test.txt"))
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

func ReadMainInputByte(dir string) [][]byte {
	res := [][]byte{}

	file, err := os.Open(filepath.Join(dir, "main.txt"))
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
