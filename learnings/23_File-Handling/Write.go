package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// // how to write  in the file
// func main() {
// 	f, err := os.Create("example.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	// // this method appends the Strings into the text file
// 	// f.WriteString("demo testing")
// 	// f.WriteString("Welcome to knoxverse")

// 	// another method
// 	// bytes := []byte("Hello Golang")
// 	// f.Write(bytes)

// 	// now here we are reading from one file and inserting it into another file

// 	//read and write to another file (Streaming fashion)

// 	sourceFile, err := os.Open("source.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer sourceFile.Close()

// 	destFile, err := os.Create("destination.txt")
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer destFile.Close()

// 	reader := bufio.NewReader(sourceFile)
// 	writer := bufio.NewWriter(destFile)

// 	for {
// 		b, err := reader.ReadByte()
// 		if err != nil {
// 			if err.Error() != "EOF" {
// 				panic(err)
// 			}
// 			break
// 		}
// 		e := writer.WriteByte(b)
// 		if err != nil {
// 			panic(e)
// 		}
// 	}

// 	writer.Flush()

// 	fmt.Println("The data successfully recieved in destination file.....!")

// }
