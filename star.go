package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
)

func main() {
//const MAX int = 5

// reader := bufio.NewReader(os.Stdin)
// text, _ := reader.ReadString('\n')
groupstar := ""

// j, _ := strconv.Atoi(text)
j := 6
// fmt.Println(j)
for i:=1; i < j+1; i++ {
    // fmt.Println("a")
    groupstar = ""
    for k:=1; k <= (j*2)-1; k++ {
        // fmt.Println(k)
        
        if j-i >= k  {
            groupstar = groupstar + " "
         
        } else if (j*2-1)-(j-i) < k {
            groupstar = groupstar + " "
        } else {
            groupstar = groupstar + "*"
        }
        
    }

	fmt.Println(groupstar)
}

// fmt.Println("a")

}