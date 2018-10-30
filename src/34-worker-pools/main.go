package main

import (
    "fmt"
    "time"
)

const num_worker = 3
const num_job = 5

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 0; w < num_worker; w++ {
        go worker(w, jobs, results)
    }

    for j := 0; j < num_job; j++ {
        jobs <- j
    }

    close(jobs)

    for i := 0; i < num_job; i++ {
        <-results
    }
}

