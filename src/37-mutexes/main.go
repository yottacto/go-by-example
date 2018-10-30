package main

import (
    "fmt"
    "time"
    "math/rand"
    "sync"
    "sync/atomic"
)

func main() {
    state := make(map[int]int)
    // var mutex sync.Mutex
    mutex := &sync.Mutex{}

    var readOps uint64
    var writeOps uint64

    for i := 0; i < 50; i++ {
        go func() {
            total := 0
            for {
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddUint64(&readOps, 1)

                time.Sleep(time.Nanosecond)
            }
        }()
    }

    for i := 0; i < 50; i++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
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

    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}

