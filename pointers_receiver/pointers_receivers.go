package main

import "fmt"

type BankAccount struct {
	AccountNumber string
	Balance       float64
}

func (b BankAccount) GetAccountInfo() string {

	balence := fmt.Sprint(b.Balance)
	accountInfo := "Account no : " + b.AccountNumber + " Balence : " + balence
	return accountInfo
}

func (b *BankAccount) Deposit(depositeAmount float64) {
	if b.Balance < 0 {
		return
	}
	b.Balance = b.Balance + depositeAmount
}
func (b *BankAccount) Withdraw(withDrawAmount float64) {
	if b.Balance < 0 || withDrawAmount > b.Balance || withDrawAmount < 0 {
		return
	}
	b.Balance = b.Balance - withDrawAmount
}

func main() {
	var bankAccount = BankAccount{AccountNumber: "123456789", Balance: 10}

	fmt.Println(bankAccount.GetAccountInfo())
	bankAccount.Deposit(10)
	fmt.Println("After depositing : ", bankAccount.GetAccountInfo())
	bankAccount.Withdraw(100)
	fmt.Println("After withDraw : ", bankAccount.GetAccountInfo())

}
