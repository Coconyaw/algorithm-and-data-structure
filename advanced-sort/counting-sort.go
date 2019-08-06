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

func parseStringSlice(a []int) []string {
	s := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		s[i] = strconv.Itoa(a[i])
	}
	return s
}

func countingSort(A []int) []int {
	S := make([]int, len(A))
	c := make([]int, 10000)
	for _, v := range A {
		c[v]++
	}

	total := 0
	for i := 0; i < 10000; i++ {
		total += c[i]
		c[i] = total
	}

	for i := len(A) - 1; i >= 0; i-- {
		S[c[A[i]] - 1] = A[i]
		c[A[i]]--
	}

	return S
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	A := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	S := countingSort(A)
	ans := parseStringSlice(S)
	fmt.Println(strings.Join(ans, " "))
}