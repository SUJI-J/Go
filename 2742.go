package main //시간초괴

import "fmt"

func main() {
	var row int
	fmt.Scanln(&row)
	for i := row; i > 0; i-- {
		fmt.Println(i)
	}
}
