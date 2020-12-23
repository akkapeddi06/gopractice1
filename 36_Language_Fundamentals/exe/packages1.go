package main

import (
	"fmt"
	""

)

func main() {
	fmt.Println("Hello world!")
	fmt.Println("2+3:", MySum(2, 3))
	fmt.Println("5-7", MySub(5, 7))
	fmt.Println("10-3-2", MySub(10, -3, -2))
}