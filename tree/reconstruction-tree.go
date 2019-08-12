// AOJ ALDS1_7_D: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_7_D

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var pos int
var post []int

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func index(a []int, key int) int {
	for i, v := range a {
		if v == key {
			return i
		}
	}
	return -1
}

func recursion(pre, in []int, l, r int) {
	if l >= r {
		return
	}

	// 今見てるNode
	root := pre[pos]
	pos++

	// 今見てるNodeのinOrder順での位置で
	// 左部分木と右部分木に分ける
	m := index(in, root)
	recursion(pre, in, l, m)
	recursion(pre, in, m+1, r)

	// leafまで行ったらpostorderに追加
	post = append(post, root)
}

func reConstruct(pre, in []int) {
	pos = 0
	recursion(pre, in, 0, len(pre))
	postOrder := make([]string, len(pre))
	for i := 0; i < len(pre); i++ {
		postOrder[i] = strconv.Itoa(post[i])
	}
	fmt.Println(strings.Join(postOrder, " "))
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	preOrder := make([]int, n)
	for i := 0; i < n; i++ {
		preOrder[i] = nextInt()
	}

	inOrder := make([]int, n)
	for i := 0; i < n; i++ {
		inOrder[i] = nextInt()
	}

	reConstruct(preOrder, inOrder)
}
