package main

import "fmt"

func main() {

	var num int

	fmt.Print("Введите число месяца (0-12): ")
	fmt.Scan(&num)

	switch num {
	case 1:
		fmt.Println("В этом месяце: 31 дней. ")
	case 2:
		fmt.Println("В этом месяце: 28 дней. ")
	case 3:
		fmt.Println("В этом месяце: 31 дней. ")
	case 4:
		fmt.Println("В этом месяце: 30 дней. ")
	case 5:
		fmt.Println("В этом месяце: 31 дней. ")
	case 6:
		fmt.Println("В этом месяце: 30 дней. ")
	case 7,8:
		fmt.Println("В этом месяце: 31 дней. ")
	case 9:
		fmt.Println("В этом месяце: 30 дней. ")
	case 10:
		fmt.Println("В этом месяце: 31 дней. ")
	case 11:
		fmt.Println("В этом месяце: 30 дней. ")
	case 12:
		fmt.Println("В этом месяце: 31 дней. ")
	default:
		fmt.Println("Ошибка: неизвестная операция")
	}
}
