// AOJ ALDS1_12_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_12_C

package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

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

	pq := make(PriorityQueue, 1)
	pq[0] = &Item{s, 0, 0}
	heap.Init(&pq)

	for pq.Len() > 0 {
		min := heap.Pop(&pq).(*Item)

		u := min.value
		cost := -1 * min.priority
		if g.distance[u] < cost {
			continue
		}

		g.nodes[u].color = black

		for i, v := range g.nodes[u].edge {
			if g.nodes[v].color != black {
				if g.distance[v] > g.distance[u]+g.nodes[u].weight[i] {
					g.distance[v] = g.distance[u] + g.nodes[u].weight[i]
					g.nodes[v].parent = u
					g.nodes[v].color = gray
					item := &Item{value: v, priority: -g.distance[v]}
					heap.Push(&pq, item)
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
