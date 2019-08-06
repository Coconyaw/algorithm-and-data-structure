# coding: utf-8

# AOJ ALDS1_7_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_7_A

NULL = -1

class Node():
    def __init__(self):
        self.parent =  NULL # parent node
        self.left = NULL    # first child node
        self.right = NULL   # brother node
        self.depth = NULL   # depth of node
        self.type = ""      # type of node(root, internal node, leaf)

def set_depth(T: list, r: int, d: int):
    T[r].depth = d
    if T[r].left != NULL:
        set_depth(T, T[r].left, d + 1)

    if T[r].right != NULL:
        set_depth(T, T[r].right, d)

n = int(input())

T = []
for i in range(n):
    T.append(Node())

cur = 0
for i in range(n):
    nodeinfo = [int(x) for x in input().split()]
    for j in range(nodeinfo[1]):
        if j == 0:
            T[nodeinfo[0]].left = nodeinfo[2+j]
        else:
            T[cur].right = nodeinfo[2+j]
        cur = nodeinfo[2+j]
        T[cur].parent = nodeinfo[0]

# find root node
root = 0
for i, v in enumerate(T):
    if v.parent == NULL:
        root = i
        break

# set depth to all nodes
# set_depth(T, root, 0)

# set type
for v in T:
    if v.parent == NULL:
        v.type = "root"
    elif v.left == NULL:
        v.type = "leaf"
    else:
        v.type = "internal node"

for i, v in enumerate(T):
    childlen = []
    if T[i].left != NULL:
        child = T[i].left
        childlen.append(str(child))

        while T[child].right != NULL:
            child = T[child].right
            childlen.append(str(child))

    depth = 0
    cur = i
    while T[cur].parent != NULL:
        depth += 1
        cur = T[cur].parent
            
    print("node {}: parent = {}, depth = {}, {}, {}".format(i, v.parent, depth, v.type, "[" + ", ".join(childlen) + "]"))
