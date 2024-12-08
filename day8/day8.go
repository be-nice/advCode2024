package day8

import "fmt"

type point struct {
	x, y int
}

func (p *point) add(cmp point) point {
	return point{p.x + cmp.x, p.y + cmp.y}
}

func (p *point) sub(cmp point) point {
	return point{p.x - cmp.x, p.y - cmp.y}
}

func Day8(s []string) {
	mx := make(map[point]bool)
	mxFreq := make(map[rune][]point)

	for y, row := range s {
		for x, ch := range row {
			p := point{x: x, y: y}
			mx[p] = true
			if ch != '.' {
				mxFreq[ch] = append(mxFreq[ch], p)
			}
		}
	}

	part1 := make(map[point]struct{})
	part2 := make(map[point]struct{})

	for _, tower := range mxFreq {
		for _, p1 := range tower {
			for _, p2 := range tower {
				if p1 == p2 {
					continue
				}

				pNext := p2.add(p2.sub(p1))

				if mx[pNext] {
					part1[pNext] = struct{}{}
				}

				for dir := p2.sub(p1); mx[p2]; p2 = p2.add(dir) {
					part2[p2] = struct{}{}
				}
			}
		}
	}

	fmt.Println("Part 1")
	fmt.Println(len(part1))
	fmt.Println("Part 2")
	fmt.Println(len(part2))
}
