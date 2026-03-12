 package main
import "fmt"

func processIt()func(a int)int{
	return func(a int) int{
		return 11
	}
}
func main(){

	fn:=processIt()
	fn(12)
}