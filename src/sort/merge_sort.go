package main

import (
	"fmt"
	"math/rand"
	"time"
)

func mergeSort(slice []int) []int {

	if len(slice) > 1 {
		mid := len(slice) / 2
		L := slice[:mid]
		R := slice[mid:]

		L = mergeSort(L)
		R = mergeSort(R)

		i, j, k := 0, 0, 0
		tmp := make([]int, len(slice))

		//Merge
		for i < len(L) && j < len(R) {
			if L[i] < R[j] {
				tmp[k] = L[i]
				i++
			} else {
				tmp[k] = R[j]
				j++
			}
			k++
		}

		//Copy remainig of L
		for i < len(L) {
			tmp[k] = L[i]
			i++
			k++
		}

		//Copy remaining of R
		for j < len(R) {
			tmp[k] = R[j]
			j++
			k++
		}
		return tmp
	}

	return slice
}

func main() {
	var slice = make([]int, 10)
	for index := range slice {
		rand.Seed(time.Now().UnixNano())
		var random = rand.Intn(100)
		slice[index] = random
	}
	fmt.Println("Slice before sort ", slice)
	slice = mergeSort(slice)
	fmt.Println("Slice after sort ", slice)
}
