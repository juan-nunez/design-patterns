package main
import (
    "log"
    "time"
)

type fn func()

func doWork() {
    for i := 0; i< 1000000000; i++ {

    }
}

func timingDecorator(decorating fn) fn{
    return fn(func() {
        beginTime := time.Now()
        decorating()
        since := time.Since(beginTime)
        log.Printf("Function took %s\n", since)
    })
}


func main() {
    decoratedDoWork := timingDecorator(doWork)
    decoratedDoWork()
}
