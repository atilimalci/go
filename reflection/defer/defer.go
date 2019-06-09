package main

import "fmt"

//1. A deferred function's arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	defer fmt.Println("a", i)
	i++
	return
}

//2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

//3. Deferred functions may read and assign to the returning function's named return values.
func c() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	a()
	b()
	fmt.Println("\nc returned:", c())
}
