// AOJ ALDS1_11_A: http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=ALDS1_11_A

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	adj := make([][]string, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]string, n)
		for j := 0; j < n; j++ {
			adj[i][j] = "0"
		}
	}

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		sc.Scan()
		adjList := strings.Split(sc.Text(), " ")
		outDegree, _ := strconv.Atoi(adjList[1])

		if outDegree > 0 {
			for k := 0; k < outDegree; k++ {
				j, _ := strconv.Atoi(adjList[k+2])
				j-- // convert 1 based to 0 based
				adj[i][j] = "1"
			}
		}
	}

	for _, g := range adj {
		fmt.Println(strings.Join(g, " "))
	}
}
