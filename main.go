package main

import "fmt"

func main() {

    slice := []int{1,2,3,4,5,6}

    target := 6
    found := false

    for _, value := range slice{
        if value == target{
            found = true
            break
        }
    }

    fmt.Println(found)
}
