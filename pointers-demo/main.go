package main

import "fmt"

func main() {
	a := "string"
	testPointer(&a)
	fmt.Printf("a at memory address %v  with value %s\n", &a, a)
}

func testPointer(a *string) {
	*a = "another string"
}

//For slices, arrays and maps it's not necessary to pass it as pointer to change the original value except we are manipulating the original object, like a append for example for a slice. For a map it's possible to add an element without passing the pointer.
