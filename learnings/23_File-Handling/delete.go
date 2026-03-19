package main

import (
	"fmt"
	"os"
)

func main() {

	// sourceFile, err := os.Open("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// we can also delete it using single line
	//os.Remove("example1.txt")

	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("File deleted successfully...!")
}
