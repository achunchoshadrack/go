package main

import "fmt"

func main() {
	var fname string
	fmt.Println("enter users first name: ")
	fmt.Scanln(&fname)

	var lname string
	fmt.Println("enter users last name:")
	fmt.Scanln(&lname)
	fmt.Println("your full names are:" + fname + " " + lname)
}
