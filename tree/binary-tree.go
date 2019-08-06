// AOJ ALDS1_7_B: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_7_B

package main

import (
	"fmt"
)

var NULL = -1

type node struct {
	// 親ノード, 左子ノード, 右子ノード, 子の数, 兄弟ノード, 深さ, 高さ
	parent, left, right, degree, sibling, depth, height int
}

type tree []node

func (t tree) getRoot() int {
	for i, v := range t {
		if v.parent == NULL {
			return i
		}
	}
	return NULL
}

func (t tree) setDepthAll() {
	root := t.getRoot()
	t.setDepthUnder(root, 0)
}

func (t tree) setDepthUnder(id, depth int) {
	t[id].depth = depth
	if t[id].left != NULL {
		t.setDepthUnder(t[id].left, depth+1)
	}
	if t[id].right != NULL {
		t.setDepthUnder(t[id].right, depth+1)
	}
}

func (t tree) setHeightALL() {
	root := t.getRoot()
	t.setHeightUnder(root)
}

func (t tree) setHeightUnder(id int) int {
	// height: 子ノードのHeightが大きいほう+1
	lh, rh := 0, 0
	if t[id].left != NULL {
		lh = t.setHeightUnder(t[id].left) + 1
	}
	if t[id].right != NULL {
		rh = t.setHeightUnder(t[id].right) + 1
	}

	if lh > rh {
		t[id].height = lh
	} else {
		t[id].height = rh
	}
	return t[id].height
}

func (t tree) getType(id int) string {
	if t[id].parent == NULL {
		return "root"
	}
	if t[id].left == NULL && t[id].right == NULL {
		return "leaf"
	}
	return "internal node"
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	T := make(tree, n)
	for i := 0; i < n; i++ {
		T[i] = node{NULL, NULL, NULL, 0, NULL, NULL, 0}
	}

	var id, left, right int

	// struct binary tree
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &id, &left, &right)

		if left != NULL {
			T[left].parent = id
		}
		if right != NULL {
			T[right].parent = id
		}

		if left == NULL && right == NULL {
			// 子を持たない場合
			T[id].degree = 0
		} else if left != NULL && right != NULL {
			// 2つ子を持つ場合 兄弟も設定する
			T[left].sibling = right
			T[right].sibling = left
			T[id].degree = 2
		} else {
			// 1つ子を持つ場合
			T[id].degree = 1
		}

		T[id].left = left
		T[id].right = right
	}

	T.setDepthAll()
	T.setHeightALL()
	for i, v := range T {
		fmt.Printf("node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n",
			i, v.parent, v.sibling, v.degree, v.depth, v.height, T.getType(i))
	}
}
