package main

import "fmt"

type Customer struct {
	Name, Address, Email string
	Age                  int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My Name is ", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huuuuuu from ", a.Name)
}

func main() {
	var reihan Customer
	reihan.Name = "Reihan"
	reihan.Email = "reedbulls91@gmail.com"
	reihan.Address = "Indonesia"
	reihan.Age = 16

	fmt.Println(reihan.Name)
	fmt.Println(reihan.Email)
	fmt.Println(reihan.Address)
	fmt.Println(reihan.Age)

	joko := Customer{
		Name:    "Joko",
		Address: "Cirebon",
		Age:     35,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", "reedbulls91@gmail.com", 35}
	fmt.Println(budi)

}
