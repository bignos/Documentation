// echogo prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	sep = " "
	s = os.Args[1]
	for i := 2; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}

	fmt.Println(s)
}
