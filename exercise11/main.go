package main

import "fmt"

const age int = 21

const (
	zero = iota
	one
	two
)

func main(){
	myName := "Henry"
	var input string

	fmt.Println("My name is", myName)
	fmt.Println("I am", age, "years old")
	fmt.Println("The first digit of my age is", two)
	fmt.Println("While the second digit is ", one)
	fmt.Println("This is where my \"name\" variable lives: ", &myName)
	fmt.Println("What's your name?")
	fmt.Scan(&input)
	yourName := &input
	fmt.Println("Nice to meet you,", *yourName)

}
