package tasks

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users`
}

type User struct {
	AccountNo      int            `json:"account-no"`
	HolderName     string         `json:"holder-name"`
	Pin            int            `json:"pin"`
	CurrentBalance int            `json:"current-balance"`
	Transactions   []Transactions `json:"transactions"`
}

type Transactions struct {
	Id            int    `json:"id"`
	Date          string `json:"date"`
	Amount        int    `json:"amount"`
	Type          string `json:"type"`
	BeforeBalance int    `json:"before-balance"`
	AfterBalance  int    `json:"after-balance"`
}

func Task2() {

	jsonFile, err := os.Open("bank-account.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users

	json.Unmarshal(byteValue, &users)

	var acno int
	var pin int
	fmt.Printf("Enter Account Number -> ")
	fmt.Scanln(&acno)

	fmt.Printf("Enter PIN -> ")
	fmt.Scanln(&pin)

	for i := 0; i < len(users.Users); i++ {
		if users.Users[i].AccountNo == acno && users.Users[i].Pin == pin {
			run(users, i)
			break
		} else {
			fmt.Println("Wrong Credentails")
		}
	}

	defer jsonFile.Close()

}

func run(users Users, i int) {

	var option int

	for option != 4 {
		fmt.Printf("Enter 1: Check Balance, 2: Bank Statements, 3: Withdraw, 4: Exit -> ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			fmt.Println("Your Balance is ", users.Users[i].CurrentBalance)
			break

		case 2:
			fmt.Println("Your Transactions are")
			for j := 0; j < len(users.Users[i].Transactions); j++ {
				fmt.Println("\tId: " + strconv.Itoa(users.Users[i].Transactions[j].Id))
				fmt.Println("\tDate: " + users.Users[i].Transactions[j].Date)
				fmt.Println("\tAmount: " + strconv.Itoa(users.Users[i].Transactions[j].Amount))
				fmt.Println("\tTransaction: " +
					strconv.Itoa(users.Users[i].Transactions[j].BeforeBalance) + " - " +
					strconv.Itoa(users.Users[i].Transactions[j].Amount) + " = " +
					strconv.Itoa(users.Users[i].Transactions[j].AfterBalance))
				fmt.Println("---------------------------------------")

			}
			break

		case 3:
			var amount int
			fmt.Printf("Enter amount to be withdrwan -> ")
			fmt.Scanln(&amount)

			var l int = len(users.Users[i].Transactions)

			newTransaction := Transactions{
				Id:            users.Users[i].Transactions[l-1].Id + 1,
				Date:          "10-09-2022",
				Amount:        amount,
				Type:          "ATM Withdraw",
				BeforeBalance: users.Users[i].Transactions[l-1].AfterBalance,
				AfterBalance:  users.Users[i].Transactions[l-1].AfterBalance - amount,
			}

			users.Users[i].Transactions = append(users.Users[i].Transactions, newTransaction)
			users.Users[i].CurrentBalance = newTransaction.AfterBalance

			fmt.Println("\tId: " + strconv.Itoa(users.Users[i].Transactions[l].Id))
			fmt.Println("\tDate: " + users.Users[i].Transactions[l].Date)
			fmt.Println("\tAmount: " + strconv.Itoa(users.Users[i].Transactions[l].Amount))
			fmt.Println("\tTransaction: " +
				strconv.Itoa(users.Users[i].Transactions[l].BeforeBalance) + " - " +
				strconv.Itoa(users.Users[i].Transactions[l].Amount) + " = " +
				strconv.Itoa(users.Users[i].Transactions[l].AfterBalance))
			fmt.Println("---------------------------------------")
			break
		}
	}

}
