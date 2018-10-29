package main

import "fmt"

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            job, more := <-jobs
            if more {
                fmt.Println("receivied job", job)
            } else {
                fmt.Println("receivied all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 0; j < 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}

