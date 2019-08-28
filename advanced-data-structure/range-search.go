package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var NULL = -1 >> 32
var np int // node number

type Point struct {
	id, x, y int
}

type TwoDTree struct {
	T []Node
	P []Point
}

type Node struct {
	location int
	left     int
	right    int
}

type T struct {
	T []Node
	P []int
}

func getLine(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func parseIntSlice(s []string) []int {
	a := make([]int, len(s))

	for i, v := range s {
		a[i] = parseInt(v)
	}
	return a
}

func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func newTree(P []int) *T {
	t := make([]Node, len(P))
	for i := 0; i < len(P); i++ {
		t[i] = Node{}
	}
	return &T{t, P}
}

func (t *T) make1DTree(l, r int) int {
	if !(l < r) {
		return -1
	}

	p := t.P[l:r]
	sort.Slice(p, func(i, j int) bool { return p[i] < p[j] })

	mid := (l + r) / 2
	n := np
	np++
	t.T[n].location = mid
	t.T[n].left = t.make1DTree(l, mid)
	t.T[n].right = t.make1DTree(mid+1, r)

	return n
}

func (t *T) printTree() {
	for i, v := range t.T {
		fmt.Printf("Id %d: location = %d, value = %d, left = %d, right = %d\n", i, v.location, t.P[v.location], v.left, v.right)
	}
}

func (t *T) find(v, sx, tx int) {
	x := t.P[t.T[v].location]
	if sx <= x && x <= tx {
		fmt.Println(t.P[t.T[v].location])
	}

	if t.T[v].left != -1 && sx <= x {
		t.find(t.T[v].left, sx, tx)
	}

	if t.T[v].right != -1 && x <= tx {
		t.find(t.T[v].right, sx, tx)
	}
}

func new2DTree(sc *bufio.Scanner) *TwoDTree {
	N := parseInt(getLine(sc))

	n := make([]Node, N)
	p := make([]Point, N)
	for i := 0; i < N; i++ {
		a := parseIntSlice(strings.Fields(getLine(sc)))
		p[i] = Point{i, a[0], a[1]}
		n[i] = Node{}
	}

	return &TwoDTree{n, p}
}

func (tt *TwoDTree) make2DTree(l, r, depth int) int {
	if !(l < r) {
		return NULL
	}

	p := tt.P[l:r]
	if depth%2 == 0 {
		// lからrの範囲をxについて昇順にソート
		sort.Slice(p, func(i, j int) bool { return p[i].x < p[j].x })
	} else {
		// lからrの範囲をyについて昇順にソート
		sort.Slice(p, func(i, j int) bool { return p[i].y < p[j].y })
	}

	mid := (l + r) / 2
	n := np
	np++

	tt.T[n].location = mid
	tt.T[n].left = tt.make2DTree(l, mid, depth+1)
	tt.T[n].right = tt.make2DTree(mid+1, r, depth+1)

	return n
}

func (tt *TwoDTree) find(v, sx, tx, sy, ty, depth int, ans *[]int) {
	// 見ている点のx座標y座標を取得
	x := tt.P[tt.T[v].location].x
	y := tt.P[tt.T[v].location].y

	// 範囲内なら出力
	if sx <= x && x <= tx && sy <= y && y <= ty {
		*ans = append(*ans, tt.P[tt.T[v].location].id)
	}

	if depth%2 == 0 {
		// 木の深さが偶数ならxについて比較
		if tt.T[v].left != NULL && sx <= x {
			tt.find(tt.T[v].left, sx, tx, sy, ty, depth+1, ans)
		}
		if tt.T[v].right != NULL && x <= tx {
			tt.find(tt.T[v].right, sx, tx, sy, ty, depth+1, ans)
		}
	} else {
		// 木の深さが奇数ならyについて比較
		if tt.T[v].left != NULL && sy <= y {
			tt.find(tt.T[v].left, sx, tx, sy, ty, depth+1, ans)
		}
		if tt.T[v].right != NULL && y <= ty {
			tt.find(tt.T[v].right, sx, tx, sy, ty, depth+1, ans)
		}
	}
}

func main() {

	// 1DTree
	/*
		t := newTree(A)
		np = 0
		t.make1DTree(0, N)

		t.printTree()
		t.find(0, 6, 9)
	*/
	sc := bufio.NewScanner(os.Stdin)
	tt := new2DTree(sc)
	tt.make2DTree(0, len(tt.P), 0)

	q := parseInt(getLine(sc))

	for i := 0; i < q; i++ {
		r := parseIntSlice(strings.Fields(getLine(sc)))
		var ans []int
		tt.find(0, r[0], r[1], r[2], r[3], 0, &ans)
		sort.Slice(ans, func(i, j int) bool { return ans[i] < ans[j] })
		for i := 0; i < len(ans); i++ {
			fmt.Println(ans[i])
		}
		fmt.Println()
	}
}
