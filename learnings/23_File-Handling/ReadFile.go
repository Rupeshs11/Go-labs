package main

import (
	"fmt"
	"os"
)

func main() {

	// // Read file
	// f ,err := os.Open("example.txt")
	// if err != nil {
	// 	panic (err)
	// }

	// defer f.Close()

	// buf := make([]byte , 50)

	// d,err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i:=0;i<len(buf);i++{
	// 	println("data",d,string(buf[i]))
	// }

	// //another and simplest way to read
	// but it is not best solution as it load whole file data in the memory at once

	data, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

}
