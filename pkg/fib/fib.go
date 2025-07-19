package fib

import (
	"math/rand"
	"time"
)

func FibRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return FibRecursive(n-1) + FibRecursive(n-2)
}

func FibIterative(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func FibMemoization(n int) int {
	memo := make(map[int]int)
	if n <= 1 {
		return n
	}
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}
	return memo[n]
}

func FibRandom(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func FibGoroutine(n int) int {
	if n <= 1 {
		return n
	}
	ch := make(chan int, 2)
	go func() { ch <- FibRecursive(n - 1) }()
	go func() { ch <- FibRecursive(n - 2) }()
	return <-ch + <-ch
}
