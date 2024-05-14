package main

import ("fmt";"time")

func main(){
	fmt.Print("Arryas\n")
	var intArray [5]int32
	intArray[0] = 123
	fmt.Println(intArray[0])
	fmt.Println(intArray[1:3])

	fmt.Println(&intArray[0])
	fmt.Println(&intArray[1])

	arr1 := [5]int{}
	fmt.Println(arr1)

	// '=' is used for assigning a value to a variable that has already been declared.
	// ':=' is used to declare and initialize a new variable.

	// jab var keyword aaye tab use  = aur jab nahi aaye tab ':='

	prices := [...]int{3,4,5} // Inferred Length of array means computer will assign the memory by itslef
	prices[2]= 8
	fmt.Println(prices[2])

	var intArray1 [7]int32
	intArray1[1] = 10
	fmt.Println(intArray1[1:3])


	var intSlice []int32 = []int32 {4,5,6}
	fmt.Printf("Length :%v, Capacity: %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 8)
	fmt.Printf("Length :%v, Capacity: %v\n", len(intSlice), cap(intSlice))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint8{"Adam": 23, "Jon Doe": 30}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["James"])

	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam")
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	for name:= range myMap2{
		fmt.Printf("Name: %v\n", name)
	}

	for i:=0; i<10;i++{
		fmt.Println(i)
	}

	var n int = 100000
	var testSlice = []int {}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice)<n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}