package main

import "fmt"

func checkAge(age int) string{
	if age < 18{
		return "Вы несовершенолетний"
	} else if age <= 60{
		return "Вы совершенолетний"
	} else {
		return "Вы пенсионер"
	}
}

func main(){
	var age int
	fmt.Println("Введите ваш возраст: ")
	fmt.Scan(&age)

	result := checkAge(age)
	fmt.Println(result)
}