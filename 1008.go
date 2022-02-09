package main

import "fmt"

func main() {
	var num1 int
	var num2 int

	fmt.Scanln(num1, num2)
	var result float32 = (float32)(num1 / num2)
	fmt.Printf("%f", result)
}
