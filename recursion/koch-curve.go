// AOJ ALDS1_5_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_5_C

package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y float64
}

func kochCurve(d int, p1, p2 point) {
	if d == 0 {
		return
	}

	th := float64(math.Pi * 60 / 180.0)

	// TODO p1, p2 から 頂点 s, t, u を計算
	var s, t, u point
	s.x = (2*p1.x + p2.x) / 3
	s.y = (2*p1.y + p2.y) / 3
	t.x = (p1.x + 2*p2.x) / 3
	t.y = (p1.y + 2*p2.y) / 3
	u.x = (t.x-s.x)*math.Cos(th) - (t.y-s.y)*math.Sin(th) + s.x
	u.y = (t.x-s.x)*math.Sin(th) + (t.y-s.y)*math.Cos(th) + s.y

	// first line of koch curve
	kochCurve(d-1, p1, s)

	// second
	fmt.Println(s.x, s.y)
	kochCurve(d-1, s, u)

	// third
	fmt.Println(u.x, u.y)
	kochCurve(d-1, u, t)

	// forth
	fmt.Println(t.x, t.y)
	kochCurve(d-1, t, p2)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	p1 := point{0, 0}
	p2 := point{100, 0}

	fmt.Println(p1.x, p1.y)
	kochCurve(n, p1, p2)
	fmt.Println(p2.x, p2.y)
}
