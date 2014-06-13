package main

import (
    "fmt"
    "flag"
    "log"
    "os"
)

var logType = flag.String("logtype", "stderr", "-logtype [stderr|syslog]")

func main() {
    flag.Parse()
    
    switch {
        case *logType == "stderr":
            fmt.Println("Logging using stderr")
        case *logType == "syslog":
            fmt.Println("Logging using syslog")
        default:
            fmt.Println("Incorrect flag passed! Run again with -h to see flag usage")
            os.Exit(1)
    }

    log.Println("bam")
}