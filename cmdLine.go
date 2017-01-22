// Go program for command line options. command line primitives are defined in os package

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

        fmt.Println("variable argument length ", len(os.Args))
        fmt.Println(" first variable argument ", os.Args[0])
        fmt.Println(" second variable argument ", os.Args[1])
        fmt.Println(" second variable argument ", os.Args[2])

        // Slice 
        fmt.Println(" slice variable argument ", os.Args[1:])

        /* For loop */
	for i := len(os.Args); i > 0; i-- {
		s += sep + os.Args[i-1]
		sep = " "
	}

	fmt.Println(s)
}
