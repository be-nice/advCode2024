package day16

import (
	"container/heap"
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
)

var (
	dx = []int{0, 1, 0, -1}
	dy = []int{1, 0, -1, 0}
)

type Point struct {
	x, y int
}

type Node struct {
	Point
	dir int
}

type State struct {
	Node
	cost  int
	index int
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*State)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func Day16(r [][]rune) {
	start, end := findStart(r)
	cost, par := dijkstra(r, Node{start, 0})
	par = cleanPars(par)

	ans := -1

	endNodes := []Node{}

	for key, val := range cost {
		if key.Point == end {
			if ans == -1 || val < ans {
				ans = val
				endNodes = []Node{key}
			} else if val == ans {
				endNodes = append(endNodes, key)
			}
		}
	}

	visitedNodes := getVisitedNodes(par, endNodes)

	fmt.Println("Part 1")
	fmt.Println(ans)
	fmt.Println("Part 2")
	fmt.Println(len(visitedNodes))
}

func nMod(curr, val int) int {
	res := curr % val

	if res < 0 {
		res += val
	}

	return res
}

func dijkstra(grid [][]rune, s Node) (map[Node]int, map[Node][]Node) {
	cost := map[Node]int{}
	marked := map[Node]bool{}
	par := map[Node][]Node{}
	cost[s] = 0
	pq := make(PriorityQueue, 0)

	heap.Init(&pq)

	item := &State{
		s,
		0,
		0,
	}

	heap.Push(&pq, item)

	for pq.Len() > 0 {
		curr := heap.Pop(&pq).(*State)

		marked[curr.Node] = true

		adj := []State{}

		adj = append(adj, State{
			Node: Node{
				Point{curr.x, curr.y},
				nMod(curr.Node.dir+1, len(dx)),
			},
			cost: cost[curr.Node] + 1000,
		})

		adj = append(adj, State{
			Node: Node{
				Point{curr.x, curr.y},
				nMod(curr.Node.dir-1, len(dx)),
			},
			cost: cost[curr.Node] + 1000,
		})

		px := curr.x + dx[curr.dir]
		py := curr.y + dy[curr.dir]

		if grid[px][py] != '#' {
			adj = append(adj, State{
				Node: Node{
					Point{px, py},
					curr.Node.dir,
				},
				cost: cost[curr.Node] + 1,
			})
		}

		for _, val := range adj {
			vCost, exists := cost[val.Node]
			newCost := val.cost

			if !exists || newCost <= vCost {
				if !exists {
					par[val.Node] = []Node{curr.Node}
				} else {
					if newCost == vCost {
						par[val.Node] = append(par[val.Node], curr.Node)
					} else {
						par[val.Node] = []Node{curr.Node}
					}
				}

				cost[val.Node] = newCost
				vState := &State{
					Node: val.Node,
					cost: newCost,
				}

				heap.Push(&pq, vState)
			}
		}
	}

	return cost, par
}

func cleanPars(par map[Node][]Node) map[Node][]Node {
	res := map[Node][]Node{}

	for key, val := range par {
		uniq := make(map[Node]struct{})

		for _, it := range val {
			uniq[it] = struct{}{}
		}

		mapSlice := make([]Node, 0, len(uniq))

		for k := range uniq {
			mapSlice = append(mapSlice, k)
		}

		res[key] = mapSlice
	}

	return res
}

func getVisitedNodes(par map[Node][]Node, st []Node) []Point {
	queue := []Node{}

	res := mapset.NewSet[Point]()

	queue = append(queue, st...)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		res.Add(curr.Point)

		p, exists := par[curr]

		if exists {
			for _, it := range p {
				queue = append(queue, it)
			}
		}
	}

	return res.ToSlice()
}

func findStart(r [][]rune) (Point, Point) {
	s := Point{}
	e := Point{}

	for i, row := range r {
		for j, ch := range row {
			if ch == 'S' {
				s.x = i
				s.y = j
			}

			if ch == 'E' {
				e.x = i
				e.y = j
			}
		}
	}

	return s, e
}
