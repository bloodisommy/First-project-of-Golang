package main

import "fmt"

func main() {

    myStudents := make(map[string]int)

    myStudents["Ruben"] = 19
    myStudents["Anatoly"] = 10
    myStudents["Miqo"] = 82

    var oldestStudent string
    var maxAge int

   for name, age := range myStudents{
    if age > maxAge{
        maxAge = age
        oldestStudent = name
     }
   }

   fmt.Println("Самый старший студент: ", oldestStudent,",", "его возраст: ", maxAge, "лет")

}