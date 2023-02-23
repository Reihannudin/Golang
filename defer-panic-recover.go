package main

import "fmt"

// defer
func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}

// Recover
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

// panic
func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {

	//Defer
	runApplication(10)

	//panic
	runApp(true)
	fmt.Println("Eko")
}
