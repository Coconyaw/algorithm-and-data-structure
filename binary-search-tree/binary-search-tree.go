// AOJ ALDS1_8_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_8_A
// AOJ ALDS1_8_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_8_B
// AOJ ALDS1_8_C: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_8_C

package main

import (
	"bufio"
	"fmt"
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

func find(key int) *node {
	// start root node
	x := root
	for x != nil {
		// find key
		if x.key == key {
			return x
		}

		// move node
		if x.key > key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return nil
}

func delete(n *node) {

	// 削除対象のnodeをyとする
	// nodeに子がない場合または子が1つの場合はnodeを削除対象にする
	// 子が2つある場合はnodeの次節点を削除対象にする
	var y *node
	if n.left == nil || n.right == nil {
		y = n
	} else {
		y = getSuccessor(n)
	}

	// yの子xを決める
	var x *node
	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	// yの親子を繋ぎ変えてyを削除
	if x != nil {
		x.parent = y.parent // xの親を設定
	}
	if y.parent == nil {
		root = x // yがrootのとき、xをrootに設定する
	} else if y == y.parent.left {
		y.parent.left = x // yがyの親の左の子なら、xをyの親の左の子として設定する
	} else {
		y.parent.right = x // yがyの親の右の子なら、xをyの親の右の子として設定する
	}

	if y != n { // nの次節点が削除された場合
		n.key = y.key // yのkeyをnにコピーする
	}
}

func getSuccessor(n *node) *node {
	if n.right != nil {
		return getMinimum(n.right)
	}

	y := n.parent
	for y != nil && n == y.right {
		n = y
		y = n.parent
	}
	return y
}

func getMinimum(x *node) *node {
	for x.left != nil {
		x = x.left
	}
	return x
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
			if exist := find(key); exist != nil {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "delete":
			key = nextInt()
			if node := find(key); node != nil {
				delete(node)
			}
		case "print":
			printList(inorder(root, []int{}))
			printList(preorder(root, []int{}))
		}
	}
}
