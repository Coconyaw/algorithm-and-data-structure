package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

type card struct {
	suit string
	number int
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func partition(A []card, p, r int) int {
	std := A[r].number
	i := p - 1
	for j := p; j < r; j++ {
		if A[j].number <= std {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	i++
	A[i], A[r] = A[r], A[i]
	return i
}

func isStable(A, S []card) bool {
	for i := 0; i < len(A); i++ {
		if A[i] != S[i] {
			return false
		}
	}	
	return true
}

func mergeSort(A []card, left, right int) {
	if left + 1 < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid, right)
		merge(A, left, mid, right)
	}
}

func merge(A [] card, left, mid, right int) {
	n1 := mid - left
	n2 := right - mid
	L := make([]card, n1+1)
	R := make([]card, n2+1)

	infCard := card{"Z", 1 << 32}

	for i := 0; i < n1; i++ {
		L[i] = A[left+i]
	}
	L[n1] = infCard

	for i := 0; i < n2; i++ {
		R[i] = A[mid+i]
	}
	R[n2] = infCard

	i := 0 // Left array index
	j := 0 // Right array index
	for k := left; k < right; k++ {
		if L[i].number <= R[j].number {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func quickSort(A []card, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}

func main() {
	input := nextString()
	n, _ := strconv.Atoi(input)

	A := make([]card, n)
	S := make([]card, n)
	for i := 0; i < n; i++ {
		cardInfo := strings.Fields(nextString())
		number, _ := strconv.Atoi(cardInfo[1])
		c := card{cardInfo[0], number}
		A[i] = c
		S[i] = c
	}

	mergeSort(A, 0, n)
	quickSort(S, 0, len(S)-1)
	if stable := isStable(A, S); stable {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}

	for _, c := range S {
		fmt.Printf("%s %d\n", c.suit, c.number)
	}
}