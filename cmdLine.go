// Go program for command line options. command line primitives are defined in os package

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	if len(os.Args) <= 1 {
		fmt.Println("ERROR::No command line argument passed")
		return
	}

	fmt.Println("variable argument length ", len(os.Args))

	// Slice
	fmt.Println(" slice variable argument ", os.Args[1:])

	// Print arguments from end to first
	for i := len(os.Args); i > 0; i-- {
		s += sep + os.Args[i-1]
		sep = " "
	}
	fmt.Println(s)

	// make s empty string
	s = ""

	// To overcome unused variable requirement in go
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

}
