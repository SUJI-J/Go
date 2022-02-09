package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	var num3 int

	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	fmt.Println((num1 + num2) % num3)
	fmt.Println(((num1 % num3) + (num2 % num3)) % num3)

	fmt.Println((num1 * num2) % num3)
	fmt.Println(((num1 % num3) * (num2 % num3)) % num3)
}
