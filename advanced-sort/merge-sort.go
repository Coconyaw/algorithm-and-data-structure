// AOJ ALDS1_5_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_5_B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var cnt = 0

const inf = 1 << 32

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

/* mergeSort ソートアルゴリズムの１つ
*  1. 指定されたn個の要素を含む部分配列をそれぞれ n/2 個の要素を含む２つの部分配列に分割する
*  2. 分割した2つの部分配列をそれぞれmergeSortでソートする
*  3. 得られたソート済み部分配列を merge により統合する
 */
func mergeSort(A []int, left, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid, right)
		merge(A, left, mid, right)
	}
}

func merge(A []int, l, m, r int) {
	// 部分配列の要素数
	n1 := m - l
	n2 := r - m

	L := make([]int, n1+1)
	R := make([]int, n2+1)

	// もとの配列 A から部分配列を作成
	for i := 0; i < n1; i++ {
		L[i] = A[l+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = A[m+i]
	}

	// 番兵を用意
	L[n1] = inf
	R[n2] = inf

	// 分割された2スライス昇順に並べ替えながらマージ
	// O(n1 + n2)
	var i, j int
	for k := l; k < r; k++ {
		cnt++
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

func parseStringSlice(a []int) []string {
	s := make([]string, len(a))
	for i := 0; i < len(a); i++ {
		s[i] = strconv.Itoa(a[i])
	}
	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	A := make([]int, n)

	for i := 0; i < n; i++ {
		A[i] = nextInt()
	}

	mergeSort(A, 0, n)
	ans := parseStringSlice(A)
	fmt.Println(strings.Join(ans, " "))
	fmt.Println(cnt)
}
