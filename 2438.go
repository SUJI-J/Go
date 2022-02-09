package main

import "fmt"

func main() {
	var row int

	fmt.Scanf("%d", &row)

	for i := 1; i <= row; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}
