// AOJ ALDS1_4_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_4_A

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var ErrNotFound = errors.New("Linear Search: Key is not found.")

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func linearSearch(a []int, key int) (int, error) {
	for i, v := range a {
		if v == key {
			return i, nil
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
		if idx, err := linearSearch(S, num); idx != -1 && err == nil {
			contains++
		}
	}
	fmt.Println(contains)
}
