package main

import "fmt"

func main() {
	
	
// pointers
	variavel :=10

	abc(&variavel)

	fmt.Println(variavel)
}

func abc(a *int) {
	*a=200
}