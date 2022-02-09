package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Scanln(&num1, &num2)

	var result int = (num1 + num2)

	fmt.Print(result)

}
