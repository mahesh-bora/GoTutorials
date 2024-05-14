package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)
	// intDivision(6, 3)
	var numerator = 6
	var denominator = 3
	fmt.Println(intDivision(numerator, denominator))
	var result, remainder, err = intDivision(numerator, denominator)
	// ------------------------------------------
	// if err != nil{
	// 	fmt.Printf(err.Error())
	// }else if remainder == 0{
	// 	fmt.Printf("The result is %v", result)
	// }else{
	// 	//V for value
	// fmt.Printf("The result of the string is %v, the remainder is %v", result, remainder)
	// 	}
	// ------------------------------------------
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v\n", result)
	default:
		fmt.Printf("The result of the integer divison is %v with remainder %v\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {

	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	// fmt.Println(result)
	return result, remainder, err
}
