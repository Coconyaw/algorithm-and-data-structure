package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"errors"
)

var sc = bufio.NewScanner(os.Stdin)
var ErrNotFound = errors.New("Linear Search: Key is not found.")

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func binarySearch(a []int, key int) (int, error) {
	var left, right, mid int
	right = len(a)
	for left < right {
		mid = (left + right) / 2
		if a[mid] == key {
			return mid, nil
		}
		if a[mid] > key {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1, ErrNotFound
}

func main() {
	sc.Split(bufio.ScanWords)
	n, _ := strconv.Atoi(nextString())
	S := make([]int, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(nextString())
		S[i] = num
	}

	q, _ := strconv.Atoi(nextString())
	var contains int
	for i := 0; i < q; i++ {
		num, _ := strconv.Atoi(nextString())
		if idx, err := binarySearch(S, num); idx != -1 && err == nil {
			contains++
		}
	}
	fmt.Println(contains)
}