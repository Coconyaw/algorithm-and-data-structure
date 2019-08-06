package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func display(A []int, N int) {
	array := make([]string, N)
	for i := 0; i < N; i++ {
		array[i] = strconv.Itoa(A[i])
	}
	fmt.Println(strings.Join(array, " "))
}

// InsertionSort Sort int slice by insertion sort algorythm
// average of quontity of calculation is O(N^2)
func InsertionSort(A []int, N int) {
	display(A, N)
	for i := 1; i < N; i++ {
		v := A[i]
		j := i - 1
		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = v
		display(A, N)
	}
}

func main() {
	line := nextLine()
	length, _ := strconv.Atoi(line)

	line = nextLine()
	array := make([]int, length)
	splitedLine := strings.Split(line, " ")

	for i := 0; i < length; i++ {
		num, _ := strconv.Atoi(splitedLine[i])
		array[i] = num
	}

	InsertionSort(array, len(array))
}
