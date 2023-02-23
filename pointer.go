package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}
func main() {
	var address1 Address = Address{"Tangerang", "Banten", "Indonesia"}

	var address2 *Address = &address1
	var address3 *Address = &address1

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)

	eko := Man{"Eko"}
	eko.Married()

	fmt.Println(eko.Name)
}
