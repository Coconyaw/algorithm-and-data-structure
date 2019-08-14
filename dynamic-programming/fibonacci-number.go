// AOJ ALDS1_10_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_10_A

package main

import (
	"fmt"
)

func fibonacci(n int) int64 {
	f := []int64{1, 1}
	for i := 2; i <= n; i++ {
		f = append(f, f[i-2]+f[i-1])
	}
	return f[n]
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Println(fibonacci(n))
}
