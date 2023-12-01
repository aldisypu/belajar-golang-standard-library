package main

import "fmt"

func main() {
	firstName := "Aldi"
	lastName := "Putra"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
