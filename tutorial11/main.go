//Generics

package main

import (
	"fmt"
)

func main() {
	var intSlice = []int{}
	fmt.Println(sumSlice[int](intSlice))
	fmt.Println(isEmpty(intSlice))

	var floatSlice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](floatSlice))
	fmt.Println(isEmpty(floatSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}

func isEmpty[T any](slice []T) bool{
	return len(slice)==0
}