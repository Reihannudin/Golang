package main

import "fmt"

// function with parameter
func SayHello(name string) string {
	//fmt.Println("Hello ", name)
	//	return value
	return "Hello " + name
}

func Bio(firstname string, lastname string) string {
	//fmt.Println("Hello ", name)
	//	return multiple value
	return "Hello " + firstname + " " + lastname
}

// misses return value
func GetFullName() (string, string) {
	return "Kim", "Seokjin"
}

// named Return value
func GetCompleteName() (firstName, middleName, lastName string) {
	firstName = "Louis"
	middleName = "William"
	lastName = "Tomlinson"

	return firstName, middleName, lastName
}

// variadic function
func SumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

// function value
func GetGoodBye(name string) string {
	return "Good Bye " + name + "See You In Future"
}

// function as parameter

// with function type declarations
type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string2 string) string) {
func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}
func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

// Annonymous Function
type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

//Recursive Function

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i // result = result * i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	result1 := SayHello("Reihan")
	fmt.Println(result1)

	result2 := Bio("Andrian", "Raihannudin")
	fmt.Println(result2)

	firstName, _ := GetFullName()
	fmt.Println(firstName)

	firstName, middleName, lastName := GetCompleteName()
	fmt.Println(firstName, middleName, lastName)

	total1 := SumAll(60, 34, 16)
	fmt.Println(total1)

	goodbye := GetGoodBye
	fmt.Println(goodbye("Nadia"))

	sayHelloWithFilter("Jungkook", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)

	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("eko", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("eko", func(name string) bool {
		return name == "root"
	})

	loop := factorialLoop(10)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1)

	recursive := factorialRecursive(10)
	fmt.Println(recursive)
}
