// AOJ ALDS1_11_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_11_C

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var distance []int

type queue struct {
	q    []int
	head int
	tail int
}

func (q *queue) push(n int) {
	if cap(q.q)-1 <= q.tail {
		q.q = append(q.q, n)
	} else {
		q.q[q.tail] = n
	}
	q.tail++
}

func (q *queue) pop() int {
	q.head++
	return q.q[q.head-1]
}

func (q *queue) isEmpty() bool {
	return q.head == q.tail
}

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

func bfs(graph [][]int, s int) {
	distance = make([]int, len(graph))

	infty := -1
	for i := 0; i < len(distance); i++ {
		distance[i] = infty
	}
	distance[s] = 0

	q := &queue{}
	q.push(s)

	for !q.isEmpty() {
		u := q.pop()
		for i := 0; i < len(graph); i++ {
			if graph[u][i] == 0 {
				continue // u, i間に辺がない場合
			}
			if distance[i] != infty {
				continue // 訪問済みの場合
			}
			// 訪問するnodeの始点からの距離を設定し、訪問キューに入れる
			distance[i] = distance[u] + 1
			q.push(i)
		}
	}

	// print distance
	for i, d := range distance {
		fmt.Printf("%d %d\n", i+1, d)
	}
}

func main() {
	adj := makeAdjList()
	bfs(adj, 0)
}
