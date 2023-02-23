package main

import "fmt"

func main() {

	var name string

	name = "reihan"
	fmt.Println(name)

	username := "Louis"
	fmt.Println(username)

	var (
		firstName = "Andrian"
		lastName  = "Raihannudin"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	const crush = "Nadia"
	fmt.Println(crush)

	//cannot replace bcs constant variabel
	//crush := "Maya"

	type NoID string

	var NoIDReihan NoID = "13240703613"
	fmt.Println(NoIDReihan)

}
