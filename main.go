package main

import "fmt"

func main() {

    const (
        Cold = iota * 10-5
        Warm
        Hot
    )

    fmt.Println(Cold, Warm, Hot)

}
