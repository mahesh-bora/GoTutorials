package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var a int = 5
	// var a int16 = 5
	var b bool = true
	//unsigned int stores only positive values
	// var c uint = 10
	var f float32 = 100.5
	var c uint64 = 10
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Number is:", a)
	fmt.Println("Boolean is:", b)
	fmt.Println(c)
	fmt.Println(f)

	var myString string
	myString = "Hello This is a String"
	fmt.Println(myString)
	fmt.Println(len(myString))

	var demoString string
	demoString = "$α"
	fmt.Println(demoString)
	fmt.Println(len(demoString))

	var length int = utf8.RuneCountInString(demoString)
	fmt.Println("This is the length from Rune Count In String:", length)

	var myRune rune = 'A'
	fmt.Println(myRune) // Runes are weird

	var1, var2 := 1, 2
	fmt.Println(var1)
	fmt.Println(var1 + var2)

	const pi = 3.14
	const myConst string = "const value"
	fmt.Println(myConst)
}
