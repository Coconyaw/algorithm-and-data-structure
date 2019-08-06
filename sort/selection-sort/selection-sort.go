// AOJ ALDS1_2_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_2_B

package main

import (
	"fmt"
)

func trace(A []int, N int) {
	for i := 0; i < N; i++ {
		if i > 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", A[i])
	}
	fmt.Printf("\n")
}

func selectionSort(A []int, N int) {
	var exchangeCount int
	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		if i != minj {
			A[i], A[minj] = A[minj], A[i]
			exchangeCount++
		}
	}
	trace(A, N)
	fmt.Println(exchangeCount)
}

func main() {
	var length int
	fmt.Scanf("%d", &length)

	array := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scanf("%d", &array[i])
	}
	selectionSort(array, length)
}
