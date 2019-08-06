package main

/*
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Node struct {
	key string
	prev *Node
	next *Node
}

type LinkedList struct {
	topnode *Node
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

func initList() *LinkedList {
	list := &LinkedList{&Node{"", nil, nil}}
	list.topnode.prev = list.topnode
	list.topnode.next = list.topnode
	return list
}

func (l *LinkedList) insert(key string) {
	node := &Node{key, l.topnode, nil}
	node.next = l.topnode.next
	l.topnode.next.prev = node
	l.topnode.next = node
}

func (l *LinkedList) delete(key string) {
	node := l.searchList(key)
	if node != nil {
		l.deleteNode(node)	
	}
}

func (l *LinkedList) deleteNode(n *Node) {
	if n == l.topnode {
		return 
	}
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *LinkedList) deleteFirst() {
	l.deleteNode(l.topnode.next)
}

func (l *LinkedList) deleteLast() {
	l.deleteNode(l.topnode.prev)
}

func (l *LinkedList) printList() {
	curNode := l.topnode
	count := 0
	for curNode.next != l.topnode {
		if count != 0 {
			fmt.Printf(" ")
		}
		curNode = curNode.next
		fmt.Printf("%s", curNode.key)
		count++
	}
	fmt.Printf("\n")
}

func (l *LinkedList) searchList(key string) *Node {
	// find first node which have the key.
	curNode := l.topnode
	for curNode.next != l.topnode {
		curNode = curNode.next
		if curNode.key == key {
			return curNode
		}
	}
	return nil
}

func main() {
	sc.Split(bufio.ScanWords)

	inputLines := nextInt()
	list := initList()

	for i := 0; i < inputLines; i++ {
		command := nextString()

		switch command {
		case "insert":
			key := nextString()
			list.insert(key)
		case "delete":
			key := nextString()
			list.delete(key)
		case "deleteFirst":
			list.deleteFirst()
		case "deleteLast":
			list.deleteLast()
		default:
			fmt.Println("Unknown Command.")
		}
	}

	list.printList()
}
*/
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"container/list"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

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
	sc.Split(bufio.ScanWords)

	inputLines := nextInt()
	linkList := list.New()

	for i := 0; i < inputLines; i++ {
		command := nextString()

		switch command {
		case "insert":
			key := nextString()
			
			linkList.PushFront(key)
		case "delete":
			key := nextString()
			for e := linkList.Front(); e != nil; e = e.Next() {
				if v := e.Value.(string); v == key {
					linkList.Remove(e)	
					break
				}
			}
		case "deleteFirst":
			linkList.Remove(linkList.Front())
		case "deleteLast":
			linkList.Remove(linkList.Back())
		default:
			fmt.Println("Unknown Command.")
		}
	}

	array := make([]string, linkList.Len())
	e := linkList.Front()
	for i := 0; i < linkList.Len(); i++ {
		array[i] = e.Value.(string)
		e = e.Next()
	}

	fmt.Println(strings.Join(array, " "))
}