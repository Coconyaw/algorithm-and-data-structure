// AOJ DSL_1_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=DSL_1_A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type DisjointSet struct {
	rank []int
	P    []int
}

func newDisjointSet(size int) *DisjointSet {
	rank := make([]int, size)
	P := make([]int, size)
	dj := &DisjointSet{rank, P}
	dj.initDisjointSet()
	return dj
}

func (dj *DisjointSet) initDisjointSet() {
	for i := 0; i < len(dj.P); i++ {
		dj.makeSet(i)
	}
}

func (dj *DisjointSet) makeSet(x int) {
	dj.P[x] = x
	dj.rank[x] = 0
}

func (dj *DisjointSet) same(x, y int) bool {
	return dj.findSet(x) == dj.findSet(y)
}

func (dj *DisjointSet) findSet(x int) int {
	if x != dj.P[x] {
		dj.P[x] = dj.findSet(dj.P[x])
	}
	return dj.P[x]
}

func (dj *DisjointSet) unite(x, y int) {
	dj.link(dj.findSet(x), dj.findSet(y))
}

func (dj *DisjointSet) link(x, y int) {
	if dj.rank[x] > dj.rank[y] {
		dj.P[y] = x
	} else {
		dj.P[x] = y
		if dj.rank[x] == dj.rank[y] {
			dj.rank[y]++
		}
	}
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	q := nextInt(sc)

	dj := newDisjointSet(n)

	for i := 0; i < q; i++ {
		c := nextInt(sc)
		x := nextInt(sc)
		y := nextInt(sc)

		if c == 0 {
			dj.unite(x, y)
		} else {
			if same := dj.same(x, y); same {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	}
}
