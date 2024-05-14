package main

import ("fmt")

func main(){
	var thing1 = [5]float32{1,2,3,4,5}
	fmt.Printf("\nThe memory locations of the thing1 array is: %p", &thing1)
	var result [5]float32 = square(&thing1)
	fmt.Printf("\nThr result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}

func square(thing2 *[5]float32) [5]float32{
	fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2)
	for i := range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return *thing2
}