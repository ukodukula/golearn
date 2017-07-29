package main

import "fmt"

func main() {

    var maxFib int = 10

    for i := 0; i <= maxFib; i++ {
        fmt.Printf("fib(%v) = %v\n", i, fib(i))
    }
}

func fib(num int) int {

    if num < 0 {
        panic ("Negatives not supported.")
    }

    if num == 1 || num == 0 {
        return num	
    } else {
        return fib(num - 1) + fib(num - 2)
    }
}
