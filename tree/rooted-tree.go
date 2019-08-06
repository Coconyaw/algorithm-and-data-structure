// AOJ ALDS1_7_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_7_A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var NULL = -1

type node struct {
	parent, left, right, depth int
}

func nextInt() int {
	sc.Scan()
	v, _ := strconv.Atoi(sc.Text())
	return v
}

// 再帰的に深さを求める
func recursion(nodes []node, n, p int) {
	nodes[n].depth = p
	if nodes[n].right != NULL {
		recursion(nodes, nodes[n].right, p)
	}

	if nodes[n].left != NULL {
		recursion(nodes, nodes[n].left, p+1)
	}
}

func getNodeType(n node) string {
	if n.parent == NULL {
		return "root"
	}

	if n.left == NULL {
		return "leaf"
	}

	return "internal node"
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()

	nodes := make([]node, n)
	for i := 0; i < n; i++ {
		nodes[i] = node{NULL, NULL, NULL, NULL}
	}

	// structing tree
	var l int
	for i := 0; i < n; i++ {
		id := nextInt()
		k := nextInt()
		for j := 0; j < k; j++ {
			c := nextInt()
			if j == 0 {
				nodes[id].left = c
			} else {
				nodes[l].right = c
			}
			l = c
			nodes[c].parent = id
		}
	}

	// find root node
	var root int
	for i := 0; i < n; i++ {
		if nodes[i].parent == NULL {
			root = i
			break
		}
	}

	recursion(nodes, root, 0)
	for i, v := range nodes {
		var childlen []string
		for i := nodes[i].left; i != NULL; i = nodes[i].right {
			childlen = append(childlen, strconv.Itoa(i))
		}
		childlenStr := strings.Join(childlen, ", ")
		nodeType := getNodeType(v)
		fmt.Printf("node %d: parent = %d, depth = %d, %s, %s\n", i, v.parent, v.depth, nodeType, "["+childlenStr+"]")
	}
}
