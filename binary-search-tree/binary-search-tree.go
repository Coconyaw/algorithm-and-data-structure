// AOJ ALDS1_8_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_8_A
// AOJ ALDS1_8_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_8_B

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type node struct {
	key                 int
	parent, left, right *node
}

var root *node
var sc = bufio.NewScanner(os.Stdin)

func insert(key int) {
	var y *node
	x := root

	z := &node{key, nil, nil, nil}

	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y

	if y == nil {
		root = z
	} else {
		if z.key < y.key {
			y.left = z
		} else {
			y.right = z
		}
	}
}

func find(key int) bool {
	// start root node
	x := root
	for x != nil {
		// find key
		if x.key == key {
			return true
		}

		// move node
		if x.key > key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return false
}

func inorder(u *node, list []int) []int {
	if u != nil {
		list = inorder(u.left, list)
		// fmt.Printf(" %v", u.key)
		list = append(list, u.key)
		list = inorder(u.right, list)
	}
	return list
}

func preorder(u *node, list []int) []int {
	if u != nil {
		// fmt.Printf(" %v", u.key)
		list = append(list, u.key)
		list = preorder(u.left, list)
		list = preorder(u.right, list)
	}
	return list
}

func printList(list []int) {
	buf := bufio.NewWriter(os.Stdout)
	for _, v := range list {
		buf.Write([]byte(fmt.Sprintf(" %d", v)))
	}
	buf.Write([]byte("\n"))
	buf.Flush()
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

func main() {
	var command string
	var n, key int

	fmt.Scanf("%d", &n)

	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		command = nextString()
		switch command {
		case "insert":
			key = nextInt()
			insert(key)
		case "find":
			key = nextInt()
			if exist := find(key); exist {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "print":
			printList(inorder(root, []int{}))
			printList(preorder(root, []int{}))
		}
	}
}