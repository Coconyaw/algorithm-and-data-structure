// AOJ GRL_1_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=GRL_1_C

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Graph [][]int

var INF = math.MaxInt32

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func (g Graph) floyd() {
	for k := 0; k < len(g); k++ {
		for i := 0; i < len(g); i++ {
			if g[i][k] == INF {
				continue
			}
			for j := 0; j < len(g); j++ {
				if g[k][j] == INF {
					continue
				}
				g[i][j] = min(g[i][j], g[i][k]+g[k][j])
			}
		}
	}
}

func (g Graph) printAPSP() {
	negative := false

	// 負の閉路がないかチェック
	for i := 0; i < len(g); i++ {
		if g[i][i] < 0 {
			negative = true
		}
	}

	if negative {
		fmt.Println("NEGATIVE CYCLE")
		return
	}

	for _, v := range g {
		ans := ""
		for _, d := range v {
			if d == INF {
				ans += fmt.Sprintf("INF ")
			} else {
				ans += fmt.Sprintf("%d ", d)
			}
		}
		ans = strings.Trim(ans, " ")
		fmt.Println(ans)
	}

}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	e := nextInt(sc)

	G := make(Graph, n)
	for i := 0; i < n; i++ {
		g := make([]int, n)
		for j := 0; j < n; j++ {
			g[j] = INF
			if i == j {
				g[j] = 0
			}
		}
		G[i] = g
	}

	var s, t, d int
	for i := 0; i < e; i++ {
		s = nextInt(sc)
		t = nextInt(sc)
		d = nextInt(sc)
		G[s][t] = d
	}
	G.floyd()
	G.printAPSP()
}
