// AOJ ALDS1_10_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_10_B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)

	P := make([]int, n+1)
	for i := 1; i <= n; i++ {
		l := nextInt(sc)
		P[i-1] = l
		m := nextInt(sc)
		P[i] = m
	}

	matrix := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		matrix[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			matrix[i][j] = (1 << 21)
			for k := i; k <= j-1; k++ {
				matrix[i][j] = min(matrix[i][j], matrix[i][k]+matrix[k+1][j]+P[i-1]*P[k]*P[j])
			}
		}
	}
	fmt.Println(matrix[1][n])
}
