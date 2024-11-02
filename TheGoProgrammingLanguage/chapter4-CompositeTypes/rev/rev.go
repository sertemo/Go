package main

import "fmt"

func reverse(s []int){
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func pReverse(ptr *[]int){
	// desreferenciamos el slice
	s := *ptr
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s)
	pReverse(&s)
	fmt.Println(s)
}