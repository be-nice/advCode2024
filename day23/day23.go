package day23

import (
	"fmt"
	"slices"
	"strings"
)

func Day23(s []string) {
	graph := make(map[string]map[string]struct{})

	for _, val := range s {
		parts := strings.Split(val, "-")
		addEdge(graph, parts[0], parts[1])
	}

	maxClique := part2(graph)
	slices.Sort(maxClique)

	fmt.Println("Part 1")
	fmt.Println(part1(graph))
	fmt.Println("Part 2")
	fmt.Println(strings.Join(maxClique, ","))
}

func addEdge(graph map[string]map[string]struct{}, a, b string) {
	if graph[a] == nil {
		graph[a] = make(map[string]struct{})
	}

	if graph[b] == nil {
		graph[b] = make(map[string]struct{})
	}

	graph[a][b] = struct{}{}
	graph[b][a] = struct{}{}
}

func part1(graph map[string]map[string]struct{}) int {
	count := 0

	for a := range graph {
		for b := range graph[a] {
			if a >= b {
				continue
			}

			for c := range graph[b] {
				if b >= c {
					continue
				}

				if _, ok := graph[a][c]; !ok {
					continue
				}

				if strings.HasPrefix(a, "t") || strings.HasPrefix(b, "t") || strings.HasPrefix(c, "t") {
					count++
				}
			}
		}
	}

	return count
}

func bronKerbosch(graph map[string]map[string]struct{}, R, P, X map[string]struct{}, cliques *[][]string) {
	if len(P) == 0 && len(X) == 0 {
		clique := []string{}

		for v := range R {
			clique = append(clique, v)
		}

		*cliques = append(*cliques, clique)
		return
	}

	for v := range P {
		newR := copySet(R)
		newR[v] = struct{}{}

		newP := intersect(graph[v], P)
		newX := intersect(graph[v], X)

		bronKerbosch(graph, newR, newP, newX, cliques)

		delete(P, v)
		X[v] = struct{}{}
	}
}

func copySet(original map[string]struct{}) map[string]struct{} {
	copy := make(map[string]struct{}, len(original))

	for k := range original {
		copy[k] = struct{}{}
	}

	return copy
}

func intersect(a, b map[string]struct{}) map[string]struct{} {
	intersection := make(map[string]struct{})

	for k := range a {
		if _, ok := b[k]; ok {
			intersection[k] = struct{}{}
		}
	}

	return intersection
}

func part2(graph map[string]map[string]struct{}) []string {
	cliques := [][]string{}
	P := make(map[string]struct{})
	X := make(map[string]struct{})
	R := make(map[string]struct{})

	for v := range graph {
		P[v] = struct{}{}
	}

	bronKerbosch(graph, R, P, X, &cliques)

	maxClique := []string{}
	for _, clique := range cliques {
		if len(clique) > len(maxClique) {
			maxClique = clique
		}
	}

	return maxClique
}
