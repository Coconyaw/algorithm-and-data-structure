package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func parseStringSlice(a []int, n int) []string {
	s := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		if i == n {
			s[i] = "[" + strconv.Itoa(a[i]) + "]"
		} else {
			s[i] = strconv.Itoa(a[i])
		}
	}
	return s
}

func partition(A []int, p, r int) int {
	std := A[r-1]
	i := p - 1
	for j := p; j < r; j++ {
		if A[j] <= std {
			i++
			A[j], A[i] = A[i], A[j]
		}
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	A := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	idx := partition(A, 0, n)
	ans := parseStringSlice(A, idx)
	fmt.Println(strings.Join(ans, " "))
}