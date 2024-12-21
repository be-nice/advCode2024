package day13

import (
	"fmt"
	"strconv"
	"strings"
)

type clawMachine struct {
	btnA, btnB struct {
		x, y int
	}
	prize struct {
		c, d int
	}
}

func Day13(s []string) {
	var data []clawMachine
	var curr clawMachine
	// re := regexp.MustCompile("\\d+")

	// for _, block := range strings.Split(s, "\n\n") {
	//	var curr clawMachine
	//	matches := re.FindAllString(block, -1)
	//	x, _ := strconv.Atoi(matches[0])
	//	y, _ := strconv.Atoi(matches[1])
	//	x1, _ := strconv.Atoi(matches[2])
	//	y1, _ := strconv.Atoi(matches[3])
	//	ex, _ := strconv.Atoi(matches[4])
	//	ey, _ := strconv.Atoi(matches[5])

	//	curr.btnA.x = x
	//	curr.btnA.y = y
	//	curr.btnB.x = x1
	//	curr.btnB.y = y1
	//	curr.prize.c = ex
	//	curr.prize.d = ey

	//	data = append(data, curr)

	//}

	for _, line := range s {
		if strings.HasPrefix(line, "Button") {
			parts := strings.Split(line, " ")
			label := strings.Split(parts[1], ":")[0]
			x, _ := strconv.Atoi(parts[2][2 : len(parts[2])-1])
			y, _ := strconv.Atoi(parts[3][2:])

			if label == "A" {
				curr.btnA = struct{ x, y int }{x, y}
			} else {
				curr.btnB = struct{ x, y int }{x, y}
			}
		} else if strings.HasPrefix(line, "Prize") {
			parts := strings.Split(line, " ")
			c, _ := strconv.Atoi(parts[1][2 : len(parts[1])-1])
			d, _ := strconv.Atoi(parts[2][2:])
			curr.prize = struct{ c, d int }{c, d}

			data = append(data, curr)
		}
	}

	fmt.Println("Part 1")
	fmt.Println(solve(data, false))
	fmt.Println("Part 2")
	fmt.Println(solve(data, true))
}

func solve(data []clawMachine, part2 bool) int {
	tokens := 0
	add := 0

	if part2 {
		add = 10000000000000
	}

	for _, val := range data {
		x1, y1 := val.btnA.x, val.btnA.y
		x2, y2 := val.btnB.x, val.btnB.y
		c := val.prize.c + add
		d := val.prize.d + add

		a := float64(c*y2-d*x2) / float64(x1*y2-y1*x2)
		b := float64(d*x1-c*y1) / float64(x1*y2-y1*x2)

		if a == float64(int(a)) && b == float64(int(b)) {
			tokens += 3*int(a) + int(b)
		}
	}

	return tokens
}
