package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync/atomic"
)

type readOp struct {
    key int
    resp chan int
}

type writeOp struct {
    key int
    value int
    resp chan bool
}

func main() {
    var readOps uint64
    var writeOps uint64

    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    go func() {
        state := make(map[int]int)

        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.value
                write.resp <- true
            }
        }
    }()

    for r := 0; r < 50; r++ {
        go func() {
            for {
                read := &readOp{
                    key: rand.Intn(5),
                    resp: make(chan int),
                }
                reads <- read
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Nanosecond)
            }
        }()
    }

    for w := 0; w < 50; w++ {
        go func() {
            for {
                write := &writeOp{
                    key: rand.Intn(5),
                    value: rand.Intn(100),
                    resp: make(chan bool),
                }
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Nanosecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)

    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}

