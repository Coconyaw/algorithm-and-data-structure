// AOJ ALDS1_10_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_10_C

package main

import (
	"fmt"
)

/*
* 文字列X,Yの最長共通部分列を求める
* c[i][j]をXのi文字目, Yのj文字目までの最長共通部分列とすると
* Xi==Yjのとき c[i][j] = c[i-1][j-1j] + 1 (1文字前までの最長共通部分列+1)
* Xi!=Yjのとき c[i][j] = max(c[i-1][j], c[i][j-1]) (どちらかの文字列の1文字前までの最長共通部分列)
*/
func main() {
	var dataNum int
	var s1, s2 string

	fmt.Scanf("%d", &dataNum)

	for i := 0; i < dataNum; i++ {
		fmt.Scanf("%s", &s1)
		fmt.Scanf("%s", &s2)

		lcs := make([][]int, len(s1)+1)
		for j := 0; j < len(s1)+1; j++ {
			lcs[j] = make([]int, len(s2)+1)
		}

		for m, vs1 := range s1 {
			for n, vs2 := range s2 {
				if vs1 == vs2 {
					lcs[m+1][n+1] = lcs[m][n] + 1
				} else if lcs[m+1][n] > lcs[m][n+1] {
					lcs[m+1][n+1] = lcs[m+1][n]
				} else {
					lcs[m+1][n+1] = lcs[m][n+1]
				}
			}
		}

		// Answer
		fmt.Println(lcs[len(s1)][len(s2)])
	}
}