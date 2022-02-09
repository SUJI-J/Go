package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Scanln(&num1, &num2)

	if num1 > num2 {
		fmt.Println(">")
	} else if num1 < num2 {
		fmt.Println("<")
	} else if num1 == num2 {
		fmt.Println("==")
	}
}
