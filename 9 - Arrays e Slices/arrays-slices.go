package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arryas e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	//slice

	slice := []int{5, 4, 3, 2, 1, 0}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, -1)
	fmt.Println(slice)

	array4 := [5]string{"P1", "P2", "P3", "P4", "P5"}
	slice2 := array4[1:4]
	fmt.Println(slice2)

	array4[2] = "PÁ PUM"
	fmt.Println(slice2)
}
