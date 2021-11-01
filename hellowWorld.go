package main
 
import "fmt"
 
func main() {
    langs := [4]string{
        "Go",
        "Python",
        "Ruby",
        "PHP",
    }
 
    for i, e := range langs {
        fmt.Println(fmt.Sprintf("%d: %s", i+1, e))
    }
}

func twoSum(nums []int, target int) []int {
    var temp int
	var first1 int
	var second2 int
	for i := 0; i < len(nums); i++ {
		temp = i+1
        for j := temp ; j < len(nums); j++ {
		   
		   if nums[i] + nums[j] == target{
			first1 = i 
			second2 = j

			break
		   }
	   }
   }
   ans := [2]int{
	   nums[first1],
	   nums[second2],
   }
   return ans
    
}