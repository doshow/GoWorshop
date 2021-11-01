package main

import (
	// "bufio"
	"fmt"
	"strconv"
	// "os"
	// "strconv"
)

func main() {
	for i:=1; i<=9 ; i++{

		for k:=1; k<=9 ; k++ {
			fmt.Println(strconv.Itoa(i) + " x " + strconv.Itoa(k) + " = " + strconv.Itoa(i*k))
		}

	}

}