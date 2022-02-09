package main //배열 사용해야돼?

import "fmt"

func main() {
	var num1 int
	var num2 int
	var row int

	fmt.Scanln(&row)

	for i := 1; i <= row; i++ {
		fmt.Scan(&num1, &num2)
		fmt.Println(num1 + num2)
	}
}
