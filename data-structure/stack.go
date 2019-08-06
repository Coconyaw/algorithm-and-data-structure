package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Stack struct {
	Stack [1000]int
	Top int
}

func getStack() *Stack {
	s := &Stack{}
	return s
}

func (s *Stack) isEmpty() bool {
	return s.Top == 0
}

func (s *Stack) isFull() bool {
	return s.Top == 999
}
func (s *Stack) Push(n int) {
	s.Top++
	s.Stack[s.Top] = n
}

func (s *Stack) Pop() int {
	s.Top--
	return s.Stack[s.Top + 1]
}

func main() {
	sc.Split(bufio.ScanWords)
	stack := getStack()
	for {
		if ok := sc.Scan(); !ok {
			break
		}
		token := sc.Text()
		switch token {
		case "+":
			n1, n2 := stack.Pop(), stack.Pop()
			stack.Push(n1 + n2)
		case "-":
			n1, n2 := stack.Pop(), stack.Pop()
			stack.Push(n2 - n1)
		case "*":
			n1, n2 := stack.Pop(), stack.Pop()
			stack.Push(n1 * n2)
		default:
			num, _ := strconv.Atoi(token)
			stack.Push(num)
		}
	}
	fmt.Println(stack.Pop())
}