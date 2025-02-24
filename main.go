package main

import "fmt"

func main() {

    myStudents := map[string]int{
        "Ruben": 19,
        "Anatoly": 10,
        "Miqo": 82,
    }

    var youngestStudent string
    var minAge int = 100

   for name, age := range myStudents{
    if age < minAge{
        minAge = age
        youngestStudent = name
     }
   }

   fmt.Println("Самый младший студент: ", youngestStudent,",", "его возраст: ", minAge, "лет")

}