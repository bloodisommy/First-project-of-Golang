package main

import "fmt"

func main() {

arr := [5]int{5, 10, 20, 30, 40}
fmt.Println("Элементы массива: ")

for i := 0; i < len(arr); i++{
    fmt.Println(arr[i])
}

count := 0
    for i := 0; i <len(arr); i++{
         if arr[i] > 10{
              count++
    }
}

fmt.Println("Чисел больше 10: ", count)

sum := 0
    for i := 0; i < len(arr); i++{
        sum += arr[i]
}

avg := float64(sum) / float64(len(arr))
fmt.Println("Среднее значение: ", avg)

}
