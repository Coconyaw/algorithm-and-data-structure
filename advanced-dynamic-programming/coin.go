package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	m := nextInt(sc)

	C := make([]int, m+1)
	T := make([]int, n+1)
	for i := 1; i <= m; i++ {
		C[i] = nextInt(sc)
	}

	INF := math.MaxInt32
	for i := 0; i <= n; i++ {
		T[i] = INF
	}

	T[0] = 0
	for i := 1; i <= m; i++ {
		for j := 0; j+C[i] <= n; j++ {
			T[j+C[i]] = min(T[j+C[i]], T[j]+1)
		}
	}

	fmt.Println(T[n])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
