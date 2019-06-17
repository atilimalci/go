package main

import (
	"fmt"
	"sync"
	"time"
)

var once sync.Once

func setup() {
	fmt.Println("setup complet")
}

func do() {
	once.Do(setup)
	fmt.Println("Done")
}

func main() {
	fmt.Println("Started")
	go do()
	go do()
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}
