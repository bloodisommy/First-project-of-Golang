package main

import "fmt"

type Speaker interface {
    Speak() int
}

type Person struct {
    Age int
}

func (p Person) Speak() int {
    return (p.Age)
}

func main() {
    var speaker Speaker 
    speaker = Person{Age: 17} 

    fmt.Println(speaker.Speak())
}
