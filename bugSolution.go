```go
package main

import (


    "fmt"
    "sync"
)

func main() {
    var m sync.Mutex
    var wg sync.WaitGroup
    ch := make(chan int, 10) // Buffered channel to prevent deadlock

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            m.Lock()
            ch <- i
            m.Unlock()
        }(i)
    }

    go func() {
        wg.Wait()
        close(ch)
    }()

    for i := range ch {
        fmt.Println(i)
    }
}
```