package day17

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type elfMachine struct {
	regA       int
	regB       int
	regC       int
	ip         int
	programSeq []int
	outputSeq  []int
}

func (m *elfMachine) adv(op int) {
	m.regA /= 1 << m.getCombOp(op)
	m.ip += 2
}

func (m *elfMachine) bxl(op int) {
	m.regB ^= op
	m.ip += 2
}

func (m *elfMachine) bst(op int) {
	m.regB = m.getCombOp(op) % 8
	m.ip += 2
}

func (m *elfMachine) jnz(op int) {
	if m.regA == 0 {
		m.ip += 2
		return
	}

	m.ip = op
}

func (m *elfMachine) bxc() {
	m.regB ^= m.regC
	m.ip += 2
}

func (m *elfMachine) out(op int) {
	opValue := m.getCombOp(op)
	m.outputSeq = append(m.outputSeq, opValue%8)
	m.ip += 2
}

func (m *elfMachine) bdv(op int) {
	m.regB = m.regA / (1 << m.getCombOp(op))
	m.ip += 2
}

func (m *elfMachine) cdv(op int) {
	m.regC = m.regA / (1 << m.getCombOp(op))
	m.ip += 2
}

func (m *elfMachine) getCombOp(op int) int {
	if op >= 0 && op <= 3 {
		return op
	}

	switch op {
	case 4:
		return m.regA
	case 5:
		return m.regB
	case 6:
		return m.regC
	}

	panic("not my mistake")
}

func (m *elfMachine) exec(opcode int, op int) {
	switch opcode {
	case 0:
		m.adv(op)
	case 1:
		m.bxl(op)
	case 2:
		m.bst(op)
	case 3:
		m.jnz(op)
	case 4:
		m.bxc()
	case 5:
		m.out(op)
	case 6:
		m.bdv(op)
	case 7:
		m.cdv(op)
	default:
		panic("I wrote a bug")
	}
}

func (m *elfMachine) getOutput() string {
	var builder strings.Builder

	for i, val := range m.outputSeq {
		if i > 0 {
			builder.WriteString(",")
		}
		builder.WriteString(strconv.FormatUint(uint64(val), 10))
	}

	return builder.String()
}

func (machine *elfMachine) run() {
	for machine.ip < len(machine.programSeq)-1 {
		opcode := machine.programSeq[machine.ip]
		op := machine.programSeq[machine.ip+1]
		machine.exec(opcode, op)
	}
}

func Day17(s []string) {
	regA, _ := strconv.Atoi(s[0][12:])
	regB, _ := strconv.Atoi(s[1][12:])
	regC, _ := strconv.Atoi(s[2][12:])
	programSeqStr := s[4][9:]
	programSeqStrs := strings.Split(programSeqStr, ",")
	programSeq := make([]int, len(programSeqStrs))

	for i, str := range programSeqStrs {
		programSeq[i], _ = strconv.Atoi(str)
	}

	machine := &elfMachine{
		regA:       regA,
		regB:       regB,
		regC:       regC,
		ip:         0,
		programSeq: programSeq,
	}

	machine.run()
	fmt.Println("Part 1")
	fmt.Println(machine.getOutput())
	fmt.Println("Part 2")
	fmt.Println(part2(programSeq))
}

func part2(programSeq []int) int {
	inVal := 35184372088832 // first value(8**15) that generates correct length ouput

	for {
		m := &elfMachine{
			regA:       inVal,
			programSeq: programSeq,
		}
		m.run()

		if slices.Equal(m.outputSeq, m.programSeq) {
			return inVal
		}

		if len(m.programSeq) == len(m.outputSeq) {
			for i := len(m.programSeq) - 1; i >= 0; i-- {
				if m.programSeq[i] != m.outputSeq[i] {
					inVal += int(math.Pow(8, float64(i)))
					break
				}
			}
		}
	}
}
