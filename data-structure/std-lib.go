package main

import (
	"github.com/go-stack/stack"
	"fmt"
)

func main() {
    logCaller("%+s")
    logCaller("%v   %[1]n()")
}

func logCaller(format string) {
    fmt.Printf(format+"\n", stack.Caller(1))
}