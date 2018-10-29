package main

import (
    "fmt"
    "math/rand"
)

func vals() (int, int) {
    rand.Seed(42)
    return rand.Int(), rand.Int()
}

func main() {
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    _, c := vals()
    fmt.Println(c)
}

