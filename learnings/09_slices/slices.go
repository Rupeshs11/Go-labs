package main
import "fmt"
//it is the most used construct in go and useful method also
func main(){
//uninitialized slice is nil(null)

var nums[]int
fmt.Println(nums==nil)
fmt.Println(len(nums))

// var nums = make([]int,2,5) // (datatype,lenght,capacity(initial))
// //capacity -> max number of elements can fit
// fmt.Println(cap(nums))

//function to add elements in the array
//the capacity get auto incremente if initial capacity get fulled
nums=append(nums,1)
nums=append(nums,2)

//another method to add multi elements
// nums:=[]int{2,4,5}
// fmt.Println(nums)
// fmt.Println(cap(nums))

//add elements using indexes
nums[0]=3
fmt.Println(nums)
fmt.Println(cap(nums))

}