package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type binHeap struct {
	heap []int
	size int
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func (bh binHeap) printNodes() {
	buf := bufio.NewWriter(os.Stdout)
	for i := 1; i < bh.size+1; i++ {
		buf.Write([]byte(fmt.Sprintf("node %d: ", i)))
		buf.Write([]byte(fmt.Sprintf("key = %d, ", bh.heap[i])))

		if parent := i / 2; bh.isInRange(parent) {
			buf.Write([]byte(fmt.Sprintf("parent key = %d, ", bh.heap[parent])))
		}

		if left := 2 * i; bh.isInRange(left) {
			buf.Write([]byte(fmt.Sprintf("left key = %d, ", bh.heap[left])))
		}

		if right := (2 * i) + 1; bh.isInRange(right) {
			buf.Write([]byte(fmt.Sprintf("right key = %d, ", bh.heap[right])))
		}
		buf.Write([]byte("\n"))
		buf.Flush()
	}
}

func (bh binHeap) isInRange(i int) bool {
	return 1 <= i && i <= bh.size
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	h := nextInt(sc)

	bh := binHeap{make([]int, h+1), h}

	for i := 1; i < h+1; i++ {
		bh.heap[i] = nextInt(sc)
	}

	bh.printNodes()
}
