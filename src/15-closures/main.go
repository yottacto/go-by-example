package main

import "fmt"

func inc() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := inc()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := inc()
    fmt.Println(newInts())
}

