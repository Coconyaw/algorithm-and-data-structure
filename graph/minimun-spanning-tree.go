// AOJ ALDS1_12_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_12_A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	WHITE int = iota
	GRAY
	BLACK
)

type graph struct {
	G [][]int
	C []int
	D []int
	P []int
}

var INFTY = 1 << 32

func constructGraph() *graph {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	g := make([][]int, n)
	c := make([]int, n)
	d := make([]int, n)
	p := make([]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		fields := strings.Fields(sc.Text())
		g[i] = toIntSlice(fields)
	}

	return &graph{g, c, d, p}
}

func (g *graph) initMst() {
	for i := 0; i < len(g.G); i++ {
		g.C[i] = WHITE
		g.D[i] = INFTY
	}
	g.D[0] = 0
	g.P[0] = -1
}

func (g *graph) primMst() int {
	g.initMst()

	var u int
	for {
		mincost := INFTY
		for i := 0; i < len(g.G); i++ {
			if g.C[i] != BLACK && g.D[i] < mincost {
				mincost = g.D[i]
				u = i
			}
		}

		if mincost == INFTY {
			break
		}

		g.C[u] = BLACK

		for i := 0; i < len(g.G); i++ {
			if g.C[i] != BLACK && g.existEdge(u, i) {
				if g.G[u][i] < g.D[i] {
					g.D[i] = g.G[u][i]
					g.P[i] = u
				}
			}
		}
	}

	totalCost := 0
	for _, cost := range g.D {
		totalCost += cost
	}
	return totalCost
}

func (g *graph) existEdge(u, v int) bool {
	return g.G[u][v] != INFTY && g.G[v][u] != INFTY
}

func toIntSlice(s []string) []int {
	a := make([]int, len(s))
	for i, v := range s {
		num, _ := strconv.Atoi(v)
		if num == -1 {
			num = INFTY
		}
		a[i] = num
	}
	return a
}

func main() {
	g := constructGraph()
	fmt.Println(g.primMst())
}
