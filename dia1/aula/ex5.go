package main

import "fmt"

func main() {
	a := 15 >= 15
	b := 100 < 1000
	c := 5 != 5

	fmt.Printf("o tipo de a é %T e seu valor é %v\n", a, a)
	fmt.Printf("o tipo de a é %T e seu valor é %v\n", b, b)
	fmt.Printf("o tipo de a é %T e seu valor é %v\n", c, c)
}
