package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  string
	Amount string
}

func main() {
	var contacts []contactInfo = loadJSON[contactInfo]("tutorial12/contactInfo.json")
	fmt.Printf("\n%+v", contacts)
	
	// var purchases []contactInfo = loadJSON[contactInfo]("tutorial12/purchaseInfo.json")
	// fmt.Printf("\n%+v", purchases)

}

func loadJSON[T contactInfo | purchaseInfo](filepath string) []T{
	data, _ := ioutil.ReadFile(filepath)

	var loaded = []T{}
	json.Unmarshal(data, &loaded)

	return loaded
}