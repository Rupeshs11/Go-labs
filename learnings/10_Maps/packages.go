package main

import (
	"fmt"
	"maps"
)

func main() {
	m1 := map[string]int{"price": 40, "phone": 2}
	m2 := map[string]int{"price": 40, "phone": 2}
	fmt.Println(maps.Equal(m1, m2))

	m3 := map[string]int{"price": 40, "phone": 2}
	m4 := map[string]int{"price": 80,"phone":2}
	fmt.Println(maps.Equal(m3, m4))

}
