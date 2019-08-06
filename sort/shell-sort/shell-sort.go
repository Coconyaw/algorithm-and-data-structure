// AOJ ALDS1_2_D: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_2_D

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var count int

func trace(A []int, N int) {
	for i := N - 1; i >= 0; i-- {
		if i != N-1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", A[i])
	}
	fmt.Printf("\n")
}

func traceln(A []int, N int) {
	for i := 0; i < N; i++ {
		fmt.Println(A[i])
	}
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func insertionSort(A []int, N, g int) {
	for i := g; i < N; i++ {
		tmp := A[i]
		j := i - g
		for j >= 0 && A[j] > tmp {
			A[j+g] = A[j]
			j = j - g
			count++
		}
		A[j+g] = tmp
	}
}

func ShellSort(A []int, N int, m int, G []int) []int {
	sorted := append([]int{}, A...)
	for i := m - 1; i >= 0; i-- {
		insertionSort(sorted, N, G[i])
	}
	return sorted
}

func main() {
	sc.Split(bufio.ScanWords)
	length := nextInt()

	array := make([]int, length)
	for i := 0; i < length; i++ {
		array[i] = nextInt()
	}

	// create G and m
	var G []int
	h := 1
	for h <= length {
		G = append(G, h)
		h = 3*h + 1
	}

	sorted := ShellSort(array, length, len(G), G)

	fmt.Println(len(G))
	trace(G, len(G))
	fmt.Println(count)
	traceln(sorted, length)
}
