package main 
import "fmt"

func main(){

	whoAmI := func(i interface{}){//interface{}it defines any type means i can be of any type
		switch n := i.(type){
		case int:
			fmt.Println("it is an integer")
		case string:
			fmt.Println("it is an string")
		case bool:
			fmt.Println("it is a boolean value")
		case float32:
			fmt.Println("it is a floating integer")
		default:
			fmt.Println("other",n)
		}


	}
	whoAmI(true)

}