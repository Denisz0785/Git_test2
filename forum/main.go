package main

import "fmt"

func main() {

	var arr1 [5]string
	arr2 := [5]string{"a", "b", "c", "d", "e"}
	arr1 = arr2

	for i, v := range arr1 {
		fmt.Println(v, &arr1[i])
	}
	for i, v := range arr2 {
		fmt.Println(v, &arr2[i])
	}
}
