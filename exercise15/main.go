package main

import "fmt"
import "time"

func main() {
	choiceList := make([]string, 0)
	choiceList = append(choiceList, "Slice Exercises")
	choiceList = append(choiceList, "Map Exercises")
	choiceList = append(choiceList, "New/Make Exercises")
	choiceList = append(choiceList, "Struct Exercises")

	c := Prompt(choiceList...)
	switch c{
	case 1:
		sliceFunc()
	case 2:
		mapFunc()
	case 3:
		newFunc()
	case 4:
		structFunc()
	}
}

func sliceFunc(){
	fmt.Println("Creating an int slice using shorthand notation...")
	is := []int{0, 1, 2, 3}
	fmt.Println("\t", is)
	fmt.Println()

	fmt.Println("Creating a string slice using shorthand notation...")
	ss := []string{"zero", "one", "two", "three"}
	fmt.Println("\t", ss)
	fmt.Println()

	fmt.Println("Creating an int slice using make function...")
	is2 := make([]int, 5, 10)
	fmt.Println("\t", is2)
	fmt.Println()

	fmt.Println("Appending first int slice to second int slice...")
	is2 = append(is2, is...)
	fmt.Println("\t", is2)
	fmt.Println()

	fmt.Println("Deleting elements from second int slice...")
	is2 = append(is2[:0], is2[5:]...)
	fmt.Println("\t", is2)
	fmt.Println()

	fmt.Println("Throwing index out of range error...")
	defer checkRecover()
	is3 := make([]int, 0, 4)
	is3 = append(is3, is...)
	fmt.Println("\t", "is3 elements:", is3)
	fmt.Println("\t", "Attempting to print non-existant element 4...")
	fmt.Println("\t", is3[4])
}

func mapFunc(){
	fmt.Println("Creating a map of with string keys and string entries...")
	im := map[string]string{
		"Eddard" : "Stark",
		"Robert" : "Baratheon",
		"Ramsay" : "Snow",
	}
	printMap(im)

	fmt.Println("Adding an entry to the map...")
	im["Hoster"] = "Tully"
	printMap(im)

	fmt.Println("Changing an entry in the map...")
	im["Ramsay"] = "Bolton"
	printMap(im)

	fmt.Println("Deleting an entry from the map...")
	delete(im, "Eddard")
	printMap(im)

	fmt.Println("Using \"comma ok\" idiom to check for the key \"Tywin\"...")
	if val, ok := im["Tywin"]; ok{
		fmt.Printf("\tTywin is of the house %v.\n", val)
	} else {
		fmt.Println("\tTywin is not in the map.")
	}


}

func newFunc() {
	fmt.Println("Creating a new int variable...")
	var v1 *int = new(int)
	fmt.Println("Printing memory address of int variable...")
	fmt.Println("\t", v1)
	fmt.Println("Printing value of int variable...")
	fmt.Println("\t", *v1)
	fmt.Println("New returned a pointer.")
	fmt.Println()

	fmt.Println("\tCreating a new string variable...")
	var v2 *string = new(string)
	fmt.Println("Printing memory address of string variable...")
	fmt.Println("\t", v2)
	fmt.Println("Printing value of string variable...")
	fmt.Println("\t", *v2)
	fmt.Println("New returned a pointer.")
	fmt.Println()

	fmt.Println("Creating a new bool variable...")
	var v3 *bool = new(bool)
	fmt.Println("Printing memory address of bool variable...")
	fmt.Println("\t", v3)
	fmt.Println("Prnting value of bool variable")
	fmt.Println("\t", *v3)
	fmt.Println("New returned a pointer.")
	fmt.Println()

	fmt.Println("Making a slice of ints...")
	v4 := make([]int, 0)
	fmt.Println("Printing value of int slice...")
	fmt.Println("\t", v4)
	fmt.Println("Make did not return a pointer.")
	fmt.Println()

	fmt.Println("Creating a map of type [int]string...")
	v5 := make(map[int]string)
	fmt.Println("Printing value of map...")
	for key, val := range v5 {
		fmt.Println("/t", key, "-", val)
	}
	fmt.Println()
}

func structFunc() {
	fmt.Println("Creating customer info struct...")
	type customer struct {
		name string
		membership bool
	}

	fmt.Println("Initializing two variables of type customer...")
	c1 := customer{"Roose", true}
	c2 := customer{"Ramsay", false}

	fmt.Println("Printing using dot notation...")
	fmt.Println("\t", c1.name, "-", c1.membership)
	fmt.Println("\t", c2.name, "-", c2.membership)

	fmt.Println("Changing value of a field...")
	c2.membership = true
	fmt.Println("Printing changed field...")
	fmt.Println("\t", c2.name, "-", c2.membership)

}

func checkRecover() {
	r := recover()
	if r != nil {
		fmt.Println("\t", `Recovered from "index out of error" panic.`)
		fmt.Println()
	}
}

func printMap(im map[string]string) {
	for key, val := range im {
		if val != "Snow" {
			fmt.Println("\t", key, "of house", val)
		} else {
			fmt.Println("\t", key, "a bastard of the name", val)
		}
	}
	fmt.Println("Make did not return a pointer.")
	fmt.Println()
}

// ---------- These are just helper functions I made to make it easier on myself ----------

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