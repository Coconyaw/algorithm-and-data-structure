// AOJ ALDS1_12_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_12_B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	white int = iota
	gray
	black
)

type node struct {
	id     int
	parent int
	edge   []int
	weight []int
	color  int
}

type graph struct {
	nodes    []node
	distance []int
}

var infty = 1 << 32
var null = -1

func constructgraph() *graph {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)

	g := &graph{make([]node, n), make([]int, n)}

	for i := 0; i < n; i++ {
		id := nextInt(sc)
		edges := nextInt(sc)
		n := node{}
		n.id = id
		n.edge = make([]int, edges)
		n.weight = make([]int, edges)
		n.color = white
		for j := 0; j < edges; j++ {
			to := nextInt(sc)
			w := nextInt(sc)
			n.edge[j] = to
			n.weight[j] = w
		}
		g.nodes[id] = n
	}
	return g
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func (g *graph) preparedijkstra(s int) {
	for i := 0; i < len(g.nodes); i++ {
		g.nodes[i].color = white
		g.nodes[i].parent = null
		g.distance[i] = infty
	}
	g.distance[s] = 0
}

func (g *graph) dijkstra(s int) []int {
	g.preparedijkstra(s)

	for {
		var u int
		mincost := infty
		for i, d := range g.distance {
			if g.nodes[i].color != black && mincost > d {
				mincost = d
				u = i
			}
		}

		if mincost == infty {
			break
		}

		g.nodes[u].color = black

		for i, e := range g.nodes[u].edge {
			if g.nodes[e].color != black {
				if g.distance[e] > g.distance[u]+g.nodes[u].weight[i] {
					g.distance[e] = g.distance[u] + g.nodes[u].weight[i]
					g.nodes[e].parent = u
					g.nodes[e].color = gray
				}
			}
		}
	}
	return g.distance
}

func printShortestPath(d []int) {
	for i, v := range d {
		fmt.Printf("%d %d\n", i, v)
	}
}

func main() {
	g := constructgraph()
	d := g.dijkstra(0)
	printShortestPath(d)
}
