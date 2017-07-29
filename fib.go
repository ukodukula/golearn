package main

import "fmt"

func main() {
    for i := 0; i<100; i++ {
        fmt.Printf("fib(%v) = %v\n", i, fib(i))
    }
}

func fib(num int) int {
    if num == 1 || num == 0 {
        return num	
    } else {
        return fib(num - 1) + fib(num - 2)
    }
}
