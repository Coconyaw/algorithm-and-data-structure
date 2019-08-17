// AOJ: ALDS1_11_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_11_B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// repr state of visited node
// WHITE yet visit
// GEAY  visiting
// BLACK visited
const (
	WHITE int = iota
	GRAY
	BLACK
)

type dfsGraph struct {
	g      [][]int // adjecent list
	colors []int   // visit state
	stime  []int   // Start time of visiting node
	ftime  []int   // Finish time of visiting node
}

var time int

func makeAdjList() [][]int {
	var n int
	fmt.Scanf("%d", &n)

	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
	}

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		nodeInfo := strings.Split(sc.Text(), " ")
		outDegree, _ := strconv.Atoi(nodeInfo[1])

		if outDegree > 0 {
			for k := 0; k < outDegree; k++ {
				j, _ := strconv.Atoi(nodeInfo[k+2])
				j-- // convert 1 based to 0 based
				adj[i][j] = 1
			}
		}
	}

	return adj
}

func (dg dfsGraph) dfsVisit(u int) {
	// 頂点uを訪問中に設定
	dg.colors[u] = GRAY
	time++
	dg.stime[u] = time

	// uの隣接nodeを訪問
	for v := 0; v < len(dg.g); v++ {
		if dg.g[u][v] == 0 {
			continue // 辺がなければスキップ
		}

		// vが未訪問なら訪問
		if dg.colors[v] == WHITE {
			dg.dfsVisit(v)
		}
	}

	// 頂点uを訪問済みに設定
	dg.colors[u] = BLACK
	time++
	dg.ftime[u] = time
}

func (dg dfsGraph) dfs() {
	for i := 0; i < len(dg.g); i++ {
		if dg.colors[i] == WHITE {
			dg.dfsVisit(i)
		}
	}
}

func (dg dfsGraph) printDfsTime() {
	for i := 0; i < len(dg.g); i++ {
		fmt.Printf("%d %d %d\n", i+1, dg.stime[i], dg.ftime[i])
	}
}

func main() {
	adj := makeAdjList()
	colors := make([]int, len(adj))
	stime := make([]int, len(adj))
	ftime := make([]int, len(adj))
	graph := dfsGraph{adj, colors, stime, ftime}
	graph.dfs()
	graph.printDfsTime()
}
