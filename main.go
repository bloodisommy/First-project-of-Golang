package main

import "fmt"

type Person struct{
    Name string
    Age int
}


func main() {

    var p Person

    fmt.Print("Введите имя: ")
    fmt.Scanln(&p.Name)

    fmt.Print("Введите возраст: ")
    fmt.Scanln(&p.Age)

    fmt.Printf("Имя: %s\nВозраст: %d\n", p.Name, p.Age)

}
