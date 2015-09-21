/*
Author: Henry Sarabia
Description:
	This is a small program that will evaluate a simple mathematical expression
	using slices, scans, and recursion. The program is based on a context free grammar
	for solving infix expressions.
 */
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"time"
)

func main(){
	//Prompting the user for input
	fmt.Println("Please enter a mathematical expression with the following properties:")
	fmt.Println("1. The expression uses infix notation")
	fmt.Println("2. Included numbers must only be integers")
	fmt.Println("3. Operations must be addition, subtraction, multiplication, or division")
	fmt.Println("4. Every character must be seperated by a space")
	fmt.Println("\nAn example is as follows: 2 + 3 * 4 / 2 + 3 * 2 - 2\n")
	fmt.Print("Now enter the expression: ")

	//Creating reader and scanning input
	expReader := bufio.NewReader(os.Stdin)
	x, err := expReader.ReadString('\n')
	if err != nil {
		fmt.Println("READ STRING")
		panic(err)
	}

	//Separating characters into a slice of strings
	xSlice := strings.Fields(x)

	//Appending end of line characters to not go out of bounds during runtime
	xSlice = append(xSlice, "\n")
	xSlice = append(xSlice, "\n")

	//This counter will keep track of our "iterator's" position through the slice
	counter := 0

	//Call Exp() with the necesarry arguments to keep track of our "iterator"
	fmt.Printf("The result is %v\n", Exp(&counter, xSlice))

	//This is literally just to finish the first exercise involving variadic functions.
	//That being said, it's utterly useless to the rest of the program.
	butWaitTheresMore()

	fmt.Println("Please type like 3 numbers, press enter, and I'll add them together or something. ")
	like3Numbers := make([]int, 6)
	//I added more room for numbers because someone is bound to try adding more than three numbers
	fmt.Scanln(&like3Numbers[0], &like3Numbers[1], &like3Numbers[2], &like3Numbers[3], &like3Numbers[4], &like3Numbers[5])
	fmt.Println("That's like",stillMore(like3Numbers...), ":)\n")
}

//This function acts as the expression that only returns a non-terminal Term by calling
//the function Exp2 and the Term function as its argument.
func Exp(counter *int, xSlice []string) int{
	return Exp2(counter, xSlice, Term(counter, xSlice))
}

//This function parses through input stream stored in our slice and will perform either
//addition of subtraction on the terms based on the previous input characters.
func Exp2(counter *int, xSlice []string, inp int) int{
	result := inp
	a := xSlice[*counter]
	*counter = *counter + 1

	if a != "\n" {
		if a == "+" {
			result = Exp2(counter, xSlice, result + Term(counter, xSlice))
		} else if a == "-" {
			result = Exp2(counter, xSlice, result - Term(counter, xSlice))
		}
	}
	return result
}

//This function returns the Term2 function output while using the simple Fact function
//as its argument. This is the same as the first non-terminal T in the
//Context Free Grammar that corresponds to this program.
func Term(counter *int, xSlice[]string) int{
	return Term2(counter, xSlice, Fact(counter, xSlice))
}

//This function parses through the input stream stored in our slice and will perform
//either multiplication or division based on the previous input characters. If it reads
//in a plus or minus character, it will move back the iterator on the input stream.
func Term2(counter *int, xSlice[] string, inp int) int{
	result := inp
	a := xSlice[*counter]
	*counter = *counter + 1

	if a != "\n" {
		if a == "*" {
			result = Term2(counter, xSlice, result * Fact(counter, xSlice))
		} else if a == "/" {
			result = Term2(counter, xSlice, result / Fact(counter, xSlice))
		} else if a == "+" || a == "-" {
			*counter = *counter - 1
		}
	}
	return result
}

//This function simply returns the next character in the stream as an integer
//while iterating through our slice by 1.
func Fact(counter *int, xSlice[]string) int{
	temp := xSlice[*counter]
	*counter = *counter + 1
	temp2, err := strconv.Atoi(temp)
	if err != nil {
		fmt.Println("FACT FUNCTION")
		panic(err)
	}
	return temp2
}

//More!
func butWaitTheresMore() {
	for i := 0; i < 4; i++ {
		time.Sleep(600 * time.Millisecond)
		switch i{
		case 0:
			fmt.Print("BUT ")
		case 1:
			fmt.Print("WAIT ")
		case 2:
			fmt.Print("THERE'S ")
		case 3:
			fmt.Print("MORE!")
		}
	}
	time.Sleep(600 * time.Millisecond)
	fmt.Println()
}

//More!?
func stillMore(numbers ...int) (sum int){
	for _, value := range numbers {
		sum += value
	}
	return
}