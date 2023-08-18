package main

import (
	"fmt"
	mystrings "go-notes/mystrings"
)

func main() {
	errorExample()
	interfaceExample()
	structsExample()
	mapsExample()

	fmt.Println(mystrings.Reverse("Hello"))
}
