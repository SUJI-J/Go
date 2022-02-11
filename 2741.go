package main //시간초과 -> fmt말고 bufio로 읽어야 돼

import (
	"bufio"
	"fmt"
	"os"
)

var reade *bufio.Reader = bufio.NewReader(os.Stdin)
var write *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)
	for i := 1; i <= T; i++ {
		fmt.Fprintln(writer, i)
	}

}
