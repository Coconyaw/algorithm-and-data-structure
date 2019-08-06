// AOJ ALDS1_4_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_4_C

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

func h1(key, m int) int {
	return key % m
}

func h2(key, m int) int {
	return 1 + (key % (m - 1))
}
func hash(key, i, m int) int {
	return (h1(key, m) + i*h2(key, m)) % m
}

func hashInseart(a []int, key int) int {
	i := 0
	m := len(a)
	for {
		idx := hash(key, i, m)
		if a[idx] == 0 {
			a[idx] = key
			return idx
		}
		i++
	}
}

func hashSearch(a []int, key int) (int, error) {
	i := 0
	m := len(a)
	for {
		idx := hash(key, i, m)
		if a[idx] == key {
			return idx, nil
		}
		if a[idx] == 0 || i >= m {
			break
		}
		i++
	}
	return -1, ErrNotFound
}

func main() {
	sc.Split(bufio.ScanWords)
	n, _ := strconv.Atoi(nextString())
	S := make([]int, n)
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(nextString())
		hashInseart(S, num)
	}

	q, _ := strconv.Atoi(nextString())
	var contains int
	for i := 0; i < q; i++ {
		num, _ := strconv.Atoi(nextString())
		if idx, err := hashSearch(S, num); idx != -1 && err == nil {
			contains++
		}
	}
	fmt.Println(contains)
}
