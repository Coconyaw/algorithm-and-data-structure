// AOJ ALDS1_9_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_9_C

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type priorityQueue struct {
	queue []int
	size  int
}

var inf = math.MaxInt32

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func (pq *priorityQueue) maxHeapify(i int) {
	l := pq.left(i)
	r := pq.right(i)

	largest := i
	if l <= pq.size && pq.queue[i] < pq.queue[l] {
		largest = l
	}
	if r <= pq.size && pq.queue[largest] < pq.queue[r] {
		largest = r
	}

	if largest != i {
		pq.queue[largest], pq.queue[i] = pq.queue[i], pq.queue[largest]
		pq.maxHeapify(largest)
	}
}

func (pq *priorityQueue) left(i int) int {
	return i * 2
}

func (pq *priorityQueue) right(i int) int {
	return i*2 + 1
}

func (pq *priorityQueue) parent(i int) int {
	return i / 2
}

func (pq *priorityQueue) extract() int {
	if pq.size < 1 {
		return -inf
	}

	// rootキーが優先度max
	maxv := pq.queue[1]

	// 最後の値をrootに移しsizeを1つ減らす
	pq.queue[1] = pq.queue[pq.size]
	pq.size--

	// maxheapを保つためrootからheapifyする
	pq.maxHeapify(1)

	return maxv
}

func (pq *priorityQueue) insert(key int) {
	pq.size++
	pq.queue[pq.size] = -inf
	pq.heapIncreaseKey(pq.size, key)
}

func (pq *priorityQueue) heapIncreaseKey(i, key int) {
	if key < pq.queue[i] {
		return
	}

	pq.queue[i] = key

	for i > 1 && pq.queue[pq.parent(i)] < pq.queue[i] {
		pq.queue[pq.parent(i)], pq.queue[i] = pq.queue[i], pq.queue[pq.parent(i)]
		i = pq.parent(i)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var command string
	var key int

	pq := &priorityQueue{make([]int, 2000000), 0}

	for {
		command = nextString(sc)
		if command == "end" {
			break
		} else if command == "insert" {
			key = nextInt(sc)
			pq.insert(key)
		} else if command == "extract" {
			fmt.Println(pq.extract())
		}
	}
}
