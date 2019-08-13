// AOJ ALDS1_9_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_9_B

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type maxHeap struct {
	heap []int
	size int
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func (mh maxHeap) buildMaxHeap() {
	for i := mh.size / 2; i >= 1; i-- {
		mh.maxHeapify(i)
	}
}

func (mh maxHeap) maxHeapify(i int) {
	l := mh.left(i)
	r := mh.right(i)

	// 自分、左の子、右の子で値が最大のnodeを選ぶ
	largest := i
	if l <= mh.size && mh.heap[l] > mh.heap[i] {
		largest = l
	}

	if r <= mh.size && mh.heap[r] > mh.heap[largest] {
		largest = r
	}

	// iの子のほうが値が大きい場合
	if i != largest {
		mh.heap[i], mh.heap[largest] = mh.heap[largest], mh.heap[i]
		mh.maxHeapify(largest) // 再帰的に呼び出し
	}
}

func (mh maxHeap) left(i int) int {
	return 2 * i
}

func (mh maxHeap) right(i int) int {
	return 2*i + 1
}

func (mh maxHeap) printHeap() {
	buf := bufio.NewWriter(os.Stdout)
	for i := 1; i < mh.size+1; i++ {
		buf.Write([]byte(fmt.Sprintf(" %d", mh.heap[i])))
	}
	buf.Write([]byte("\n"))
	buf.Flush()
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	h := nextInt(sc)

	mh := maxHeap{make([]int, h+1), h}

	for i := 1; i < h+1; i++ {
		mh.heap[i] = nextInt(sc)
	}

	mh.buildMaxHeap()
	mh.printHeap()
}
