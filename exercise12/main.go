/*
Author: Henry Sarabia
Description:
	Four short functions that correspond with the four exercises on presentation 12.
	This makes use of my selector package to prompt the user for their choice in which function to run.
 */

package main

import "fmt"
import "github.com/HenrySarabia/helper/selector"

func main() {
	choiceList := make([]string, 0)
	choiceList = append(choiceList, "Mod Function")
	choiceList = append(choiceList, "Loop Function")
	choiceList = append(choiceList, "Fizz Function")
	choiceList = append(choiceList, "Natural Function")

	userChoice := selector.Prompt(choiceList...)
	switch userChoice {
	case 1:
		modFunc()
	case 2:
		loopFunc()
	case 3:
		fizzFunc()
	case 4:
		naturalFunc()
	}
}

//Function that allows the user to enter two numbers then display the remainder when one number is divided by the other
func modFunc() {
	var num1, num2 int
	fmt.Print("Please enter the first integer: ")
	fmt.Scan(&num1)
	fmt.Print("Please enter the second integer: ")
	fmt.Scan(&num2)

	fmt.Print(num1 % num2)
}

//Function that loops from 1 to 1000 whilst printing even numbers
func loopFunc() {
	loopFlag := 0
	for i := 1; i <= 1000; i += 2 {
		if loopFlag == 0 {
			i++
			loopFlag = 1
		}
		fmt.Println(i)
	}
}

//Function that prints numbers from 1 to 100 but for multiples of 3 prints "Fizz",
//multiples of 5 prints "Buzz", and for multiples of both 3 and 5 prints "FizzBuzz"
func fizzFunc() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Print("Fizz")
		}
		if i % 5 == 0 {
			fmt.Print("Buzz")
		} else if i % 3 != 0 {
			fmt.Print(i)
		}
		fmt.Println()

	}
}

//Function that finds the sum of all the multiples of 3 or 5 below 1000
func naturalFunc(){
	sum := 0
	for i := 1; i <= 1000; i++ {
		switch {
		case i % 3 == 0:
			sum += i
		case i % 5 == 0:
			sum += i
		}
	}
	fmt.Println(sum)
}