package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var inf = 1 << 32
var cnt int

func nextInt() int {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	return n
}

func mergeSort(A []int, left, right int) {
	if left + 1 < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid, right)
		merge(A, left, mid, right)
	}
}

func merge(A []int, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid

	L := make([]int, n1 + 1)
	R := make([]int, n2 + 1)

	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}

	L[n1] = inf
	R[n2] = inf

	var i, j int
	for k := left; k < right; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
			cnt += n1 - i
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	mergeSort(A, 0, n)
	fmt.Println(cnt)
}