package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        _, _ = fmt.Fprintln(os.Stderr, "usage: uuid2base64 UUID")
        os.Exit(1)
    }

    uuidStr := os.Args[1]

    result, err := ConvertToBase64(uuidStr)

    if err != nil {
        _, _ = fmt.Fprintln(os.Stderr, err)
    }

    fmt.Println(result)
}
