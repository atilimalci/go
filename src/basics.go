package main

import (
	"fmt"
	"strconv"
)

func playWithArrays() {
	fmt.Println("//Playing with arrays")
	array := [5]int{10, 20, 30, 40, 50}
	array = [...]int{10, 20, 30, 40, 50}
	array = [5]int{1: 10, 2: 20}
	fmt.Println("Array: " + strconv.Itoa(array[2]))
}

func playWithArrayPointers() {
	fmt.Println("//Playing with array pointers")
	arrayp := [5]*int{0: new(int), 1: new(int)}
	*arrayp[0] = 10
	fmt.Println("Array Pointer: ", *arrayp[0])
}

func playWithMatrix() {
	fmt.Println("//Playing with matrix")
	var matrix [4][2]int
	matrix = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println("Matrix : ", matrix[0][0])
}

func playWithSlice() {
	fmt.Println("//Playing with slice")
	slice := []int{10, 20, 30}
	fmt.Println("Slice = ", slice[1])

	stringSlice := []string{99: ""}
	fmt.Println("String Slice = ", stringSlice[0])

	// Use make to create an empty slice of integers.
	slice = make([]int, 0)
	// Use a slice literal to create an empty slice of integers.
	slice = []int{}
	fmt.Println("Slice = ", slice)

	slice = []int{10, 20, 30, 40, 50}
	// Create a new slice. Contains a length of 2 and capacity of 4 elements.
	newSlice := slice[1:3]
	fmt.Println("newSlice = ", newSlice)
	newSlice[1] = 35
	fmt.Println("After modify slice = ", slice)
	fmt.Println("After modify newSlice = ", newSlice)

	slice = []int{10, 20, 30, 40, 50}
	newSlice = slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println("After append slice = ", slice)
	fmt.Println("After append newSlice = ", newSlice)

	fmt.Println("Append-2")

	slice = []int{10, 20, 30, 40}
	newSlice = append(slice, 50)
	fmt.Println("After append slice = ", slice)
	fmt.Println("After append newSlice = ", newSlice)

	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}
}

func playWithPointers() {
	fmt.Println("//Playing with pointers")
	var a = 5.67
	var p = &a
	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("Value stored in variable p = ", p)
}

func playWithMultidimensionalSlice() {
	slice := [][]int{{10}, {100, 200}}
	fmt.Println("Multidimensional Slice ", slice)

	slice[0] = append(slice[0], 20)
	fmt.Println("Multidimensional Slice after append ", slice)
}

func playWithMap() {
	colors := map[string]string{}
	colors["Red"] = "#da1337"
	value, exists := colors["Red"]
	if exists {
		fmt.Println("Red ", value)
	}

	value = colors["Blue"]
	// Did this key exist?
	if value != "" {
		fmt.Println("Blue ", value)
	}

	delete(colors, "Red")
}

func main() {
	playWithArrays()
	playWithArrayPointers()
	playWithMatrix()
	playWithSlice()
	playWithPointers()
	playWithMultidimensionalSlice()
	playWithMap()
}
