package main

import (
    "fmt"
    "math/rand"
    "time"
)

func compute() {
    for {
        // Simulate some long running processing
        time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
        _ = fmt.Sprintf("Random number: %d", rand.Intn(1000))
    }
}

func main() {
    go compute()
    select {}
}
