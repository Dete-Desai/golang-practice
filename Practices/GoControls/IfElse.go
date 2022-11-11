package main

import "fmt"

func main() {

	var num int

	fmt.Println("Enter a number: ")
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Printf("The value is an even number\n")
	} else {
		fmt.Printf("The value is an odd number\n")
	}
	fmt.Printf("The value is: %d\n", num)
}
