package main

import (
	"fmt"
	"strings"
)

func main(){
	// fmt.Println()
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%v, %T\n", indexed, indexed)
	for i, v := range myString{
		fmt.Println(i, v)
	}

	// refer internet for runes

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v\n", myRune)

	var strSlice = [] string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i:= range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Printf("%v\n", catStr)
}