// Variables, assignments, type declarations, scope
package main

import "fmt"

func main() {

	var i, j, k int // i,j,k are variables of interger type
	var name string // name is a variable of type string 

	fmt.Printf("Uninitialized values of local variables i %d j %d k %d\n", i, j, k)
	fmt.Printf("Uninitialized values of string variable %s\n",name)
 
        name = "venkat"
	fmt.Printf("Uninitialized values of string variable %s\n",name)

        // Does go strings have NULL termination as C strings??
	fmt.Printf("length of string %d\n",len(name))
	fmt.Printf("last character of string %c\n",name[len(name)-1])

}
