package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	from := "A1"
	to := "A2"

	fmt.Printf("%s->%s = %t\n", from, to, hasEdge(from, to))
	addEdge(from, to)
	fmt.Printf("%s->%s = %t\n", from, to, hasEdge(from, to))
}
