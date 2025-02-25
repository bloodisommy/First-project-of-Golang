package main

import "fmt"

type BankAccount struct {
    Balance       float64
    AccountHolder string
}

func (b *BankAccount) Deposit(amount float64) {
    b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
    if b.Balance >= amount {
        b.Balance -= amount
    } else {
        fmt.Println("Ошибка: недостаточно средств.")
    }
}

func (b BankAccount) PrintAccountInfo() {
    fmt.Printf("Владелец счета: %s\nБаланс: %.2f\n", b.AccountHolder, b.Balance)
}

func main() {
    
    account := BankAccount{AccountHolder: "Ruben Voskanyan", Balance: 1000.50}

    account.PrintAccountInfo()

    account.Deposit(500)
    account.PrintAccountInfo()

    account.Withdraw(200)
    account.PrintAccountInfo()

    account.Withdraw(2000)
    account.PrintAccountInfo()
}
