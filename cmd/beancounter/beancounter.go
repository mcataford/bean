package main

import (
    "log"
    "time"
)

// Time between polls, in seconds.
var INTERVAL = time.Duration(10)

func main() {
    log.Println("😸 Counting beans...")
    for {
        log.Println("Collecting new datapoints.")
        time.Sleep(INTERVAL * time.Second)
    }
}
