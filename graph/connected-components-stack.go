// AOJ ALDS1_11_D: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_11_D

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

var nocolor = -1

type graph struct {
	graph  [][]int
	colors []int
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func newGraph(nodeNum int) *graph {
	colors := make([]int, nodeNum)
	g := make([][]int, nodeNum)
	for i := 0; i < nodeNum; i++ {
		g[i] = []int{}
		colors[i] = nocolor
	}

	return &graph{g, colors}
}

func constructGraph(sc *bufio.Scanner) *graph {
	usersNum := nextInt(sc)
	linkNum := nextInt(sc)

	g := newGraph(usersNum)
	for i := 0; i < linkNum; i++ {
		u := nextInt(sc) // node
		l := nextInt(sc) // link node
		// edge each other
		g.graph[u] = append(g.graph[u], l)
		g.graph[l] = append(g.graph[l], u)
	}
	return g
}

func (g *graph) dfs(r, color int) {
	g.colors[r] = color

	stack := list.New()
	stack.PushFront(r)

	for stack.Len() > 0 {
		v := stack.Remove(stack.Front()).(int)
		for _, u := range g.graph[v] {
			if g.colors[u] == nocolor {
				g.colors[u] = color
				stack.PushFront(u)
			}
		}
	}
}

func (g *graph) assignColor() {
	color := 1
	for u := 0; u < len(g.graph); u++ {
		if g.colors[u] == nocolor {
			g.dfs(u, color)
			color++
		}
	}
}

func (g *graph) isLink(s, t int) bool {
	return g.colors[s] == g.colors[t]
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	// construct graph and friendly colors
	g := constructGraph(sc)
	g.assignColor()

	// judge link
	q := nextInt(sc)
	for i := 0; i < q; i++ {
		s := nextInt(sc)
		t := nextInt(sc)

		if g.isLink(s, t) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
