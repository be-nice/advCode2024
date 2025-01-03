package day21

import (
	"container/heap"
	"fmt"
	"math"
)

var (
	seqCache  = make(map[string]int)
	pathCache = make(map[string][][]byte)
)

var ePad = [][]byte{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var rPad = [][]byte{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
}

var dirs = []Point{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

var clickMaping = map[int]byte{
	0: '<',
	1: '^',
	2: '>',
	3: 'v',
}

type Point struct {
	x, y int
}

type PriorityQueueItem struct {
	point    Point
	priority int
}

type PriorityQueue []*PriorityQueueItem

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*PriorityQueueItem)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq) - 1
	item := (*pq)[n]
	*pq = (*pq)[:n]

	return item
}

func Day21(b [][]byte) {
	run := func(depth int) int {
		res := 0

		for _, val := range b {
			cursors := make([]byte, depth+1)

			for i := range cursors {
				cursors[i] = 'A'
			}

			len := findShortestSequence(val, depth, true, cursors)

			n := 0
			for _, val := range val[:3] {
				n = n*10 + int(val-'0')
			}

			res += n * len
		}

		return res
	}

	fmt.Println("Part 1")
	fmt.Println(run(2))
	fmt.Println("Part 2")
	fmt.Println(run(25))
}

func findShortestPaths(keypad [][]byte, from, to byte) [][]byte {
	cacheKey := fmt.Sprintf("%c-%c", from, to)

	if cached, exists := pathCache[cacheKey]; exists {
		return cached
	}

	var start, end Point

	for y, row := range keypad {
		for x, ch := range row {
			if ch == from {
				start = Point{x, y}
			}

			if ch == to {
				end = Point{x, y}
			}
		}
	}

	if start == end {
		return [][]byte{{'A'}}
	}

	dist := make([][]int, len(keypad))
	visited := make([][]bool, len(keypad))
	for i := range dist {
		dist[i] = make([]int, len(keypad[i]))
		visited[i] = make([]bool, len(keypad[i]))
		for j := range dist[i] {
			dist[i][j] = math.MaxInt
		}
	}

	dist[start.y][start.x] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &PriorityQueueItem{point: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*PriorityQueueItem).point

		if visited[current.y][current.x] {
			continue
		}
		visited[current.y][current.x] = true

		for _, dir := range dirs {
			nx, ny := current.x+dir.x, current.y+dir.y

			if nx >= 0 && ny >= 0 && ny < len(keypad) && nx < len(keypad[ny]) && keypad[ny][nx] != ' ' {
				alt := dist[current.y][current.x] + 1

				if alt < dist[ny][nx] {
					dist[ny][nx] = alt
					heap.Push(pq, &PriorityQueueItem{point: Point{nx, ny}, priority: alt})
				}
			}
		}
	}

	var paths [][]byte
	var stack []struct {
		point Point
		path  []byte
	}

	stack = append(stack, struct {
		point Point
		path  []byte
	}{end, []byte{'A'}})

	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]

		if node.point == start {
			paths = append(paths, append([]byte(nil), node.path...))
			continue
		}

		for i, dir := range dirs {
			nx, ny := node.point.x+dir.x, node.point.y+dir.y
			if nx >= 0 && ny >= 0 && ny < len(keypad) && nx < len(keypad[ny]) && dist[ny][nx] < dist[node.point.y][node.point.x] {
				newPath := append([]byte{clickMaping[i]}, node.path...)
				stack = append(stack, struct {
					point Point
					path  []byte
				}{Point{nx, ny}, newPath})
			}
		}
	}

	pathCache[cacheKey] = paths

	return paths
}

func findShortestSequence(s []byte, depth int, highest bool, cursors []byte) int {
	cacheKey := fmt.Sprintf("%s:%d:%c", string(s), depth, cursors[depth])

	if val, found := seqCache[cacheKey]; found {
		return val
	}

	result := 0
	for _, ch := range s {
		var paths [][]byte
		if highest {
			paths = findShortestPaths(ePad, cursors[depth], ch)
		} else {
			paths = findShortestPaths(rPad, cursors[depth], ch)
		}

		if depth == 0 {
			minLen := math.MaxInt

			for _, path := range paths {
				if len(path) < minLen {
					minLen = len(path)
				}
			}
			result += minLen
		} else {
			minCost := math.MaxInt

			for _, path := range paths {
				cost := findShortestSequence(path, depth-1, false, cursors)

				if cost < minCost {
					minCost = cost
				}
			}

			result += minCost
		}

		cursors[depth] = ch
	}

	seqCache[cacheKey] = result

	return result
}
