package main

import "fmt"

func main() {
	sum := 0
	var num int

	fmt.Scanln(&num)

	for i := 0; i <= num; i++ {
		sum += i
	}
	fmt.Println(sum)
}
