package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("This is a C for loop.")
	}

	j := 0
	for j < 3 {
		fmt.Println("This is a C while loop.")
		j++
	}

	for {
		fmt.Println("This loop would go forever without a break statement.")
		break
	}
}