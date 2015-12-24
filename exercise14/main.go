package main

import (
	"fmt"
	"time"
	"strconv"
	"math"
)

func main() {
	choiceList := make([]string, 0)
	choiceList = append(choiceList, "Escape Function")
	choiceList = append(choiceList, "Len Function")
	choiceList = append(choiceList, "Index Function")
	choiceList = append(choiceList, "Slice Function")
	choiceList = append(choiceList, "Concatenation Function")
	choiceList = append(choiceList, "Conversions Function")
	choiceList = append(choiceList, "Rounding Up Function")

	c := Prompt(choiceList...)
	switch c{
	case 1:
		escapeFunc()
	case 2:
		lenFunc()
	case 3:
		indexFunc()
	case 4:
		sliceFunc()
	case 5:
		catFunc()
	case 6:
		convFunc()
	case 7:
		roundFunc()
	}
}

func escapeFunc() {
	fmt.Print("This is a function that uses escape sequences.\n")
	fmt.Print("\t\tSee?\n")
}

func lenFunc() {
	var input string
	fmt.Print("Please enter a string: ")
	fmt.Scan(&input)
	fmt.Println("Your string has", len(input), "bytes.")
}

func indexFunc() {
	var input  string
	fmt.Print("Please enter a string: ")
	fmt.Scan(&input)
	fmt.Print("I will now display the first rune from your string: ", input[0])
}

func sliceFunc() {
	s := []string{"h", "e", "l", "l", "o",}
	fmt.Println(s[1:])
}

func catFunc() {
	var s1 string
	var s2 string

	fmt.Print("Please enter a word: ")
	fmt.Scan(&s1)
	fmt.Print("Please enter another word: ")
	fmt.Scan(&s2)
	fmt.Println(s1+s2)
}

func convFunc() {
	var s string
	var f float64
	var b []byte

	fmt.Print("Please enter a number: ")
	fmt.Scan(&s)
	i, _ := strconv.Atoi(s)
	fmt.Println("First, let's convert this string to an integer. This is your number as an int:", i)
	s = strconv.Itoa(i)
	fmt.Println("Now, let's convert it back. This is your number as a string again:", s)
	f = float64(i)
	fmt.Println("This time we'll convert the int to a float64, here it is: ", f)
	i = int(f)
	fmt.Println("Again, let's convert it back. This is your number as an int: ", i)
	fmt.Print("Please enter a word: ")
	fmt.Scan(&s)
	b = []byte(s)
	fmt.Println("Now, let's convert it to a []byte. This is your word now: ", b)
	s = string(b)
	fmt.Println("Again, let's convert it back to a string. This is it: ", s)
}

func roundFunc() {
	var x float64

	fmt.Print("Please enter a float64: ")
	fmt.Scan(&x)
	fmt.Println(math.Ceil(x))
}
// ------------------------------------------------------------------------------------------- //
// ------------------------------------------------------------------------------------------- //
// ------------------------------------------------------------------------------------------- //



// ------------ These are just helper functions I made to make it easier on myself ----------- //

//The input can be any number of strings which will be formatted into a numbered
//list that the user can easily read through. Next, the user is asked to enter the number
//corresponding to their choice from the list. This variable, choice, is then passed into
//the next function choose() and also passed on as its integer output.
func Prompt(options ...string) int {
	var choice int
	i := 1

	fmt.Println()
	fmt.Println("Please choose which function to run from the following choices:")
	for _, value := range options {
		fmt.Printf("%v. %v\n", i, value)
		i++
	}
	fmt.Println()
	fmt.Print("Enter the number corresponding to your choice: ")
	fmt.Scan(&choice)
	fmt.Println()

	length := len(options)

	for choice > length || choice <= 0 {
		fmt.Println("That is not an available choice.")

		if length == 1{
			fmt.Println("Your only choice is 1.")
		} else if length == 2{
			fmt.Println("Your choices are 1 or 2.")
		} else {
			fmt.Printf("Your choices are from 1 to %v.\n", length)

		}

		fmt.Print("Enter the number corresponding to your choice: ")
		fmt.Scan(&choice)
		fmt.Println()
	}
	choose(choice, options...)
	return choice
}

//This helper function is passed the user's choice along with the slice of strings denoting the user's possible choices.
//The function iterates through each option until it finds which corresponds to the user's choice and outputs this string
//along with the confirmation "You chose option %v." As a final step, wait() is called.
func choose(choice int, options ...string) {
	for key, value := range options {
		if key == choice - 1 {
			fmt.Printf("You chose option %v: \"%v\"\n", choice, value)
			wait(3, 600)
		}
	}
}

//This helper function simply pauses the program for a specified amount of ticks with a specified amount of milliseconds
//between each tick. For each tick, a simple period is output to show that the program is still running but waiting for the
//user to finish reading the previous output.
func wait(ticks int, mSeconds time.Duration) {
	for i := 0; i < ticks; i++ {
		time.Sleep(mSeconds * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(mSeconds * time.Millisecond)
	fmt.Println()
	fmt.Println()
}