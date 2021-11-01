package main

import (
	"fmt"
	// "reflect"
)

func main() {
	grString := ""
	runes := []rune("Hello, 世界")
	var tempA string
	for i := len(runes)-1; i >= 0 ; i=i-1 {
		tempA = string(runes[i])
		grString = grString + tempA
		// grString = grString + fmt.runes[i]
	}
	fmt.Printf(grString)
		// str := "HiHi"

		// var result string 
		// length := len(str) 
		// for i := 0; i < length; i++ { 
		// 	result = result + fmt.Sprintf("%c", str[length-i-1]) 
		//  }

		//  fmt.Printf(result)

}