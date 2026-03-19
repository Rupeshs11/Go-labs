package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Open("example.txt")
	if err != nil {
		// log the error
		panic(err)
	}

	fileInfo, err := f.Stat()
	if err != nil {
		// log the error
		panic(err)
	}

	fmt.Println("File name :", fileInfo.Name())
	fmt.Println("Is it a folder :", fileInfo.IsDir())
	fmt.Println("Size of the file :", fileInfo.Size())
	fmt.Println("file permission :",fileInfo.Mode())
	fmt.Println("File modified at :",fileInfo.ModTime())

}
