package main 
 import(
	"fmt"
	"time"
 )
 func main(){

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	case time.Monday, time.Tuesday ,time.Wednesday,time.Thursday,time.Friday:
		fmt.Println("it's working day")
	default:
		fmt.Println("Day doesn't exist , galat choice !")
	}
 }