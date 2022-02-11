package main //시간초과 -> fmt말고 bufio로 읽어야 돼

import (
	"bufio"
	"fmt"
	"os"
)

var read *bufio.Reader = bufio.NewReader(os.Stdin)
var writ *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var T int
	fmt.Fscan(reader, &T)
	for i := T; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
}
