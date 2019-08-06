// AOJ ALDS1_3_D: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_3_D

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	Stack [20000]interface{}
	Top   int
}

type puddle struct {
	first int
	area  int
}

func getStack() *Stack {
	s := &Stack{}
	return s
}

func (s *Stack) isEmpty() bool {
	return s.Top == 0
}

func (s *Stack) isFull() bool {
	return s.Top == cap(s.Stack)
}

func (s *Stack) Push(a interface{}) {
	s.Top++
	s.Stack[s.Top] = a
}

func (s *Stack) Pop() interface{} {
	if s.isEmpty() {
		return -1
	}
	s.Top--
	return s.Stack[s.Top+1]
}

func (s *Stack) Peek() interface{} {
	if s.isEmpty() {
		return -1
	}
	return s.Stack[s.Top]
}

func main() {
	var input string
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}

	var totalArea int
	s := Stack{}
	puddles := Stack{}
	for cur, v := range input {
		switch string(v) {
		case "\\":
			s.Push(cur)
		case "/":
			p := s.Pop().(int)
			if p == -1 {
				break
			}
			curArea := cur - p
			totalArea = totalArea + curArea
			for !puddles.isEmpty() && puddles.Peek().(puddle).first > p {
				tmp := puddles.Pop().(puddle)
				curArea = curArea + tmp.area
			}
			puddles.Push(puddle{p, curArea})
		}
	}

	puddleNum := puddles.Top
	array := make([]string, puddles.Top+1)
	for {
		p := puddles.Pop()
		if p == -1 {
			break
		}
		array = append(array, strconv.Itoa(p.(puddle).area))
	}
	array = append(array, strconv.Itoa(puddleNum))

	// reverse array
	for i := len(array)/2 - 1; i >= 0; i-- {
		opp := len(array) - 1 - i
		array[i], array[opp] = array[opp], array[i]
	}

	fmt.Println(totalArea)
	fmt.Println(strings.Trim(strings.Join(array, " "), " "))
}
