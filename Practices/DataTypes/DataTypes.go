package main

import (
	"fmt"
	"math/cmplx"
)

func main() {

	var age int
	var name string
	var statement bool
	var height float64
	var intelligence complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Printf("%T %T %T %T %Tv%T %T %T %T\n", age, name, statement, height, intelligence)
	fmt.Printf("%v %q %v %v %v %v %v %v %v\n", age, name, statement, height, intelligence)

	//An array is a sequence of elements of the same data type.
	var carType [5]int
	//Arrays can also be multidimensional. We can simply create them with the following format
	var multCarType [2][3]int
	//This creates a slice with zero capacity and zero length.
	var expandedCarType []int
	//Here, the slice has an initial length of 5 and has a capacity of 10.
	definedCarType := make([]int, 5, 10)
	//The capacity of a slice can be increased by using the append or a copy function.
	definedCarType = append(definedCarType, 1, 2, 3, 4)
	//Another way to increase the capacity of a slice is to use the copy function.
	definedCarType1 := make([]int, 15)
	copy(definedCarType1, definedCarType)

	fmt.Printf("Arrays\n")

	fmt.Printf("%T %T %T %T %T\n", carType, multCarType, expandedCarType, definedCarType, definedCarType1)
	fmt.Printf("%v %v %v %v %v\n", carType, multCarType, expandedCarType, definedCarType, definedCarType1)

	//We can create a sub-slice of a slice.
	// initialize a slice with 4 len and values
	var number2 = []int{1, 2, 3, 4}
	slice1 := number2[2:]
	slice2 := number2[:3]
	slice3 := number2[1:4]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	//Maps are a data type in Go, which maps keys to values.
	var m map[string]int

	//We can add keys and values easily to a map
	// adding key/value
	m["clearity"] = 2
	m["simplicity"] = 3

	// printing the values
	fmt.Printf("Map\n")
	fmt.Println(m["clearity"])
	fmt.Println(m["simplicity"])

	//One type of data type can be converted into another using type casting.
	t := 1.1
	s := int(t)

	fmt.Printf("Typecasting\n")
	fmt.Println(s)
}
