package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("David")
	generica(222)
	generica(true)
	generica(123.45)
}
