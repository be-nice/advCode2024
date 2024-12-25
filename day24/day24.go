package day24

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Gate struct {
	op  string
	w1  string
	w2  string
	out string
}

func Day24(s string) {
	inputs, gates := parse(s)
	inputs = part1(inputs, gates)

	fmt.Println("Part 1")
	fmt.Println(binToDec(inputs))
	fmt.Println("Part 2")
	fmt.Println(strings.Join(part2(gates), ","))
}

func part1(inputs map[string]int, gates []Gate) map[string]int {
	for {
		progress := false
		for _, gate := range gates {

			if _, exists := inputs[gate.out]; exists {
				continue
			}

			val1, ok1 := inputs[gate.w1]
			val2, ok2 := inputs[gate.w2]

			if ok1 && ok2 {
				var res int
				switch gate.op {
				case "AND":
					res = val1 & val2
				case "OR":
					res = val1 | val2
				case "XOR":
					res = val1 ^ val2
				}
				inputs[gate.out] = res
				progress = true
			}
		}

		if !progress {
			break
		}
	}

	return inputs
}

func part2(gates []Gate) []string {
	ops := make(map[string]Gate, len(gates))
	revOps := make(map[Gate]string, len(gates)*2)

	for _, gate := range gates {
		ops[gate.out] = gate
		revOps[Gate{op: gate.op, w1: gate.w1, w2: gate.w2}] = gate.out
		revOps[Gate{op: gate.op, w1: gate.w2, w2: gate.w1}] = gate.out
	}

	top := 0
	re := regexp.MustCompile(`z(\d+)`)

	for key := range ops {
		if matches := re.FindStringSubmatch(key); matches != nil {
			n, _ := strconv.Atoi(matches[1])
			if n > top {
				top = n
			}
		}
	}

	wrongGates := make([]string, 0, 8)

	for i := 1; i < top; i++ {
		x := fmt.Sprintf("x%02d", i)
		y := fmt.Sprintf("y%02d", i)
		z := fmt.Sprintf("z%02d", i)

		resOp, ok := ops[z]
		if !ok {
			continue
		}

		xorGate, xorExists := revOps[Gate{op: "XOR", w1: x, w2: y}]
		andGate, andExists := revOps[Gate{op: "AND", w1: x, w2: y}]

		if !xorExists || !andExists {
			continue
		}

		if resOp.op != "XOR" {
			wrongGates = append(wrongGates, z)
		}

		carry := []string{}
		for _, op := range ops {
			if op.op == "XOR" && (op.w1 == xorGate || op.w2 == xorGate) {
				other := op.w1
				if op.w1 == xorGate {
					other = op.w2
				}
				carry = append(carry, other)
			}
		}

		if len(carry) != 1 {
			wrongGates = append(wrongGates, xorGate)
			wrongGates = append(wrongGates, andGate)
		} else {
			carryGate := carry[0]
			xor2Gate, xor2Exists := revOps[Gate{op: "XOR", w1: xorGate, w2: carryGate}]

			if xor2Exists && xor2Gate != z {
				wrongGates = append(wrongGates, xor2Gate)
			}
		}
	}

	sort.Strings(wrongGates)

	return wrongGates
}

func binToDec(inputs map[string]int) int {
	zRegex := regexp.MustCompile(`^z(\d+)$`)
	zWires := []struct {
		index string
		value int
	}{}

	for wire, value := range inputs {
		if matches := zRegex.FindStringSubmatch(wire); matches != nil {
			zWires = append(zWires, struct {
				index string
				value int
			}{index: wire, value: value})
		}
	}

	sort.Slice(zWires, func(i, j int) bool {
		return zWires[i].index < zWires[j].index
	})

	var sb strings.Builder

	for _, zWire := range zWires {
		sb.WriteString(strconv.Itoa(zWire.value))
	}

	binStr := sb.String()
	reverseBinary := func(s string) string {
		runes := []rune(s)

		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}

		return string(runes)
	}

	binStr = reverseBinary(binStr)
	res, _ := strconv.ParseInt(binStr, 2, 64)

	return int(res)
}

func parse(s string) (map[string]int, []Gate) {
	inputs := make(map[string]int)
	gates := make([]Gate, 0, len(s))

	split := strings.Split(s, "\n\n")

	for _, line := range strings.Split(strings.TrimSpace(split[0]), "\n") {
		parts := strings.Split(line, ": ")
		wire := parts[0]
		val, _ := strconv.Atoi(parts[1])
		inputs[wire] = val
	}

	for _, line := range strings.Split(strings.TrimSpace(split[1]), "\n") {
		tokens := strings.Fields(line)
		gates = append(gates, Gate{
			op:  tokens[1],
			w1:  tokens[0],
			w2:  tokens[2],
			out: tokens[4],
		})

	}

	return inputs, gates
}
