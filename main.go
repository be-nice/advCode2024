package main

import (
	"adv/day1"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	day1.Day1(ReadTestInput("day1"))
	day1.Day1(ReadMainInput("day1"))
}

func ReadTestInput(dir string) []string {
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

func ReadMainInput(dir string) []string {
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
