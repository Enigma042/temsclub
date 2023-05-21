package main

import (
	"fmt"
	"time"
)

var firstname string
var lastname string
var email string
var number int
var billvalue int
var product string
var no int
var finalbill int
var list = make([]User, 0)

func main() {
	for {
		billvalue = 0
		greetings()
		UserDet()
		finalbill = eval()
		go printbill(firstname, lastname, email, uint(finalbill))
	}

}
func greetings() {
	fmt.Println("Welcome to our store")
}
func UserDet() {

	fmt.Println("Enter your first name")
	fmt.Scanln(&firstname)
	fmt.Println("Enter your last name")
	fmt.Scanln(&lastname)
	fmt.Println("Enter your email")
	fmt.Scanln(&email)
	fmt.Println("Enter the number of products")
	fmt.Scanln(&number)
}
func eval() int {
	for i := 0; i < number; i++ {
		fmt.Println("Enter the product and the no of the product")
		fmt.Scanln(&product, &no)
		switch product {
		case "Sofa":
			billvalue = billvalue + (50000 * no)
		case "Table", "Cabinet":
			billvalue = billvalue + (15000 * no)
		case "Chair":
			billvalue = billvalue + (3000 * no)
		case "Bed":
			billvalue = billvalue + (30000 * no)
		default:
			fmt.Print("wrong input")
		}
	}
	return billvalue
}

type User struct {
	firstName string
	lastName  string
	email     string
	Bill      uint
}

func printbill(firstname string, lastname string, email string, finalbill uint) {
	time.Sleep(10 * time.Second)
	var user = User{
		firstName: firstname,
		lastName:  lastname,
		email:     email,
		Bill:      finalbill,
	}
	list = append(list, user)

	fmt.Printf("Thank you %v %v for buying goods.Your bill is %v. You will receive a confirmation email at %v\n", firstname, lastname, finalbill, email)
	fmt.Print(list)
}
