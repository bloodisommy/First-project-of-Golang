package main

import "fmt"

type Person struct{
    Name string
    Age int
}

func incrementAge(p *Person){
    p.Age = p.Age + 1
}

func main() {
  person := Person{Name: "Ruben, ", Age: 17}
  fmt.Println("До изменений: ", person)

  incrementAge(&person)

  fmt.Println("После изменений: ", person)
}