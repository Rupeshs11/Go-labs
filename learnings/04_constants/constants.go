package main

import "fmt"

const (
	port=3000    // we can declare the constants either inside or outside the main
	price="33.3" // but the shorthand method for declaration is not allowed outside
	host= "localhost"
	os="Amazon-linux 2023"
)

func main() {
	const name = "golang"

	//name := "knox"

	fmt.Println(name)
	fmt.Println(port)
	fmt.Println(price,host,os)
}
