// AOJ ALDS1_2_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_2_C

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func trace(A []string, N int) {
	for i := 0; i < N; i++ {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%s", A[i])
	}
	fmt.Printf("\n")
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func BubbleSort(A []string, N int) []string {
	bs := append([]string{}, A...)
	for i := 0; i < N; i++ {
		for j := N - 1; j > i; j-- {
			num1 := getNumberOfCard(bs[j])
			num2 := getNumberOfCard(bs[j-1])
			if num1 < num2 {
				bs[j], bs[j-1] = bs[j-1], bs[j]
			}
		}
	}
	return bs
}

func SelectionSort(A []string, N int) []string {
	ss := append([]string{}, A...)
	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			num1 := getNumberOfCard(ss[j])
			num2 := getNumberOfCard(ss[minj])
			if num1 < num2 {
				minj = j
			}
		}
		ss[i], ss[minj] = ss[minj], ss[i]
	}
	return ss
}

func getNumberOfCard(card string) int {
	num, _ := strconv.Atoi(card[1:])
	return num
}

func isStable(in, out []string, N int) {
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for a := 0; a < N; a++ {
				for b := a + 1; b < N; b++ {
					num1 := getNumberOfCard(in[i])
					num2 := getNumberOfCard(in[j])
					if num1 == num2 && in[i] == out[b] && in[j] == out[a] {
						fmt.Println("Not stable")
						return
					}
				}
			}
		}
	}
	fmt.Println("Stable")
	return
}

func main() {
	sc.Split(bufio.ScanWords)
	ns := nextString()
	length, _ := strconv.Atoi(ns)

	array := make([]string, length)
	for i := 0; i < length; i++ {
		array[i] = nextString()
	}
	bubbleSorted := BubbleSort(array, length)
	selectionSorted := SelectionSort(array, length)

	trace(bubbleSorted, length)
	isStable(array, bubbleSorted, length)
	trace(selectionSorted, length)
	isStable(array, selectionSorted, length)
}
