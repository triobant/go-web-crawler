package main

import (
    "fmt"
    "os"
)

func main() {
    argsWithoutProg := os.Args[1:]
    if len(argsWithoutProg) < 1 {
        fmt.Println("no website provided")
    } else if len(argsWithoutProg) > 1 {
        fmt.Println("too many arguments provided")
    }
    fmt.Printf("starting crawl of: %s", argsWithoutProg)
}
