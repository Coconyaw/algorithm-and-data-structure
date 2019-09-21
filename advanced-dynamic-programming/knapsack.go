// AOJ DPL_1_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=DPL_1_B

package main

import (
	"fmt"
)

type Item struct {
	v, w int
}

var Diagonal = 1
var Top = 0

func main() {
	var N, W int
	fmt.Scanf("%d %d", &N, &W)

	items := make([]Item, N+1)
	items[0] = Item{}

	var v, w int
	for i := 1; i <= N; i++ {
		fmt.Scanf("%d %d", &v, &w)
		items[i] = Item{v, w}
	}

	dp := make([][]int, N+1)
	G := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		a := make([]int, W+1)
		b := make([]int, W+1)
		dp[i] = a
		G[i] = b
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= W; j++ {
			if items[i].w <= j {
				pick := dp[i-1][j-items[i].w] + items[i].v
				notpick := dp[i-1][j]
				if pick >= notpick {
					dp[i][j] = pick
					G[i][j] = Diagonal
				} else {
					dp[i][j] = notpick
					G[i][j] = Top
				}
			} else {
				dp[i][j] = dp[i-1][j]
				G[i][j] = Top
			}
		}
	}

	fmt.Println(dp[N][W])

	/*
		// List of selected items
		selection := list.New()

		// Reconstruct selected item from G
		weight := W
		for i := N; i >= 1; i-- {
			if G[i][weight] == Diagonal {
				selection.PushFront(i)
				weight -= items[i].w
			}
		}

		// Print selected item number
		for e := selection.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value.(int))
		}
	*/
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
