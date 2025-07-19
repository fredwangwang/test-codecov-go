package main

import (
	"fmt"

	"github.com/fredwangwang/test-codecov-go/pkg/fib"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println(fib.FibRecursive(10))
	fmt.Println(fib.FibIterative(10))
	fmt.Println(fib.FibMemoization(10))
}
