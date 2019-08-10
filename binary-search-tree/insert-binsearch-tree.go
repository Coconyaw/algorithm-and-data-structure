// AOJ ALDS1_8_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_8_A

package main

import (
	"fmt"
)

type node struct {
	key int
	parent, left, right *node
}

var root *node

func insert(key int) {
	var y *node
	x := root

	// 挿入するnode
	z := &node{key, nil, nil, nil}

	// rootから木をたどってzの親を決める
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	z.parent = y

	// rootか親の左の子か右の子か決める
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

func inorder(u *node) {
	if u == nil {
		return
	}
	inorder(u.left)
	fmt.Printf(" %d", u.key)
	inorder(u.right)
}

func preorder(u *node) {
	if u == nil {
		return
	}
	fmt.Printf(" %d", u.key)
	preorder(u.left)
	preorder(u.right)
}

func main() {
	var command string
	var n, key int

	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &command)
		if command == "insert" {
			fmt.Scanf("%d", &key)
			insert(key)
		} else if command == "print" {
			inorder(root)
			fmt.Printf("\n")
			preorder(root)
			fmt.Printf("\n")
		}
	}
}