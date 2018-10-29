package main

import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return fact(n - 1) * n
}

func main() {
    fmt.Println(fact(7))
}

