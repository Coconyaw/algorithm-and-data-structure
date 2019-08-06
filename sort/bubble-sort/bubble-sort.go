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

func bubbleSort(A []int, N int) {
	exchangeCount := 0
	flag := true
	for flag {
		flag = false
		for j := N - 1; j > 0; j-- {
			if A[j] < A[j-1] {
				A[j], A[j-1] = A[j-1], A[j]
				exchangeCount++
				flag = true
			}
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
	bubbleSort(array, length)
}
