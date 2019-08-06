package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type queue struct {
	q    [50000]process
	head int
	tail int
}

type process struct {
	name string
	time int
}

type processScheduler struct {
	q        queue
	quantum  int
	execTime int
}

func (q *queue) enqueue(p process) {
	if q.isFull() {
		return
	}
	q.q[q.tail] = p
	q.tail = (q.tail + 1) % len(q.q)
}

func (q *queue) dequeue() process {
	p := q.q[q.head]
	q.head = (q.head + 1) % len(q.q)
	return p
}

func (q *queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *queue) isFull() bool {
	return q.head == (q.tail+1)%len(q.q)
}

func (ps *processScheduler) addProcess() bool {
	if ok := sc.Scan(); !ok {
		return false
	}
	pname := sc.Text()

	sc.Scan()
	ptime, _ := strconv.Atoi(sc.Text())
	ps.q.enqueue(process{pname, ptime})
	return true
}

func (ps *processScheduler) executeProcess() {

	// Get target process
	p := ps.q.dequeue()

	// Terminate Process
	if leftTime := p.time - ps.quantum; leftTime <= 0 {
		ps.execTime += p.time
		fmt.Println(p.name, ps.execTime)
		return
	}

	// re scheduling
	p.time -= ps.quantum
	ps.execTime += ps.quantum
	ps.q.enqueue(p)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	pnum, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	quantum, _ := strconv.Atoi(sc.Text())

	ps := processScheduler{queue{}, quantum, 0}

	for i := 0; i < pnum; i++ {
		ps.addProcess()
	}

	for !ps.q.isEmpty() {
		ps.executeProcess()
	}
}
