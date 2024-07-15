package main

import "fmt"

func setName(name *string) {
	*name = "b"
}
func main() {

	name := "h"

	setName(&name)

	fmt.Println(name)
}
