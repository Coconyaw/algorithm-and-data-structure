// AOJ ALDS1_5_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_5_A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func exhaustiveSearch(A []int, n, i, m int) bool {
	// i-1を選んだときにもとのmが作れている
	if m == 0 {
		return true
	}

	if i >= n {
		return false
	}

	// i を選ばなかったときに残りの数で m　を作れるか: exhaustiveSearch(A, n, i+1, m)
	// i を選んだときに残りの数で m-i　を作れるか: exhaustiveSearch(A, n, i+1, m-A[i]
	return exhaustiveSearch(A, n, i+1, m) || exhaustiveSearch(A, n, i+1, m-A[i])
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	q := nextInt()
	for i := 0; i < q; i++ {
		v := nextInt()
		if exhaustiveSearch(A, n, 0, v) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
