package main

import (
    "fmt"
    "os"
    "strconv"
    "gopl.io/ch2/popcount"
)

func main() {
    for _, args := range os.Args[1:] {
        i, err := strconv.ParseUint(args, 10, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }

        fmt.Printf("val: %v popcount: %v\n", i, popcount.PopCount(i))
    }

}
