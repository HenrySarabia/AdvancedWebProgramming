package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	choiceList := make([]string, 0)
	choiceList = append(choiceList, "Half Integer Exercise")
	choiceList = append(choiceList, "Find Max Exercise")
	choiceList = append(choiceList, "Bool Expression Exercise")
	choiceList = append(choiceList, "Two Parameter Function")
	choiceList = append(choiceList, "Two Return Function")
	choiceList = append(choiceList, "Named Return Function")

	c := Prompt(choiceList...)
	switch c{
	case 1:
		var input int
		fmt.Print("Please enter an integer: ")
		fmt.Scan(&input)
		fmt.Print("Output: ")
		fmt.Println(halfFunc(input))
	case 2:
		var input int
		is := make([]int, 0)
		fmt.Print("How many integers do you want in your list: ")
		fmt.Scan(&input)
		for i := 0; i < input; i++ {
			is = append(is, rand.Intn(1000))
		}
		fmt.Print("Integer List:")
		fmt.Println(is)
		fmt.Print("Output: ")
		fmt.Print(maxFunc(is...))
	case 3:
		boolFunc()
	case 4:
		var input string
		var input2 int
		fmt.Print("Please enter your first name: ")
		fmt.Scan(&input)
		fmt.Print("Please enter your age: ")
		fmt.Scan(&input2)
		fmt.Println(twoParamFunc(input, input2))
	case 5:
		var input string
		var input2 int
		fmt.Print("Please enter your first name: ")
		fmt.Scan(&input)
		fmt.Print("Please enter your age: ")
		fmt.Scan(&input2)
		dyears, old := twoReturnFunc(input, input2)
		fmt.Print(input, " is ", dyears, " in dog years and is ")
		if old {
			fmt.Println("old.")
		} else {
			fmt.Println("not old.")
		}
	case 6:
		var input int
		fmt.Print("Please enter your age: ")
		fmt.Scan(&input)
		fmt.Println(namedReturnFunc(input))
	}

}

func halfFunc(i int) (float64, bool) {
	var j float64 = float64(i)
	j = j / 2
	if i % 2 == 0 {
		return j, true
	} else {
		return j, false
}
}

func maxFunc(nums ...int) int{
	max := nums[0]
	for _, val := range nums{
		if max < val {
			max = val
		}
	}
	return max
}

func boolFunc() {
	fmt.Println((true && false) || (false && true) || !(false && false))
}

func twoParamFunc(name string, age int) string {
	return name + " is " + strconv.Itoa(age) + " years old."
}

func twoReturnFunc(name string, age int) (dyears int, old bool) {
	dyears = age * 7

	if age > 25 {
		old = true
	} else {
		old = false
	}
	return
}

func namedReturnFunc(age int) (dogYears int) {
	dogYears = age * 7
	return
}

func

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
			//wait(3, 600)
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