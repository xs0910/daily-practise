package test

import (
	"fmt"
	"testing"
)

func TestAddSlice(t *testing.T) {
	a := []int{0, 1, 2, 3}
	add2Slice(a, 4)
	fmt.Println(a)

	b := []int{0, 1, 2, 3}
	addSlice(&b, 4)
	fmt.Println(b)
}

func add2Slice(s []int, t int) {
	s[0]++
	fmt.Printf("slice_addr: %p, cap: %d, value: %v\n", &s[0], cap(s), s)
	s = append(s, t)
	fmt.Printf("slice_addr: %p, cap: %d, value: %v\n", &s[0], cap(s), s)
	s[0]++
	fmt.Printf("slice_addr: %p, cap: %d, value: %v\n", &s[0], cap(s), s)
}

func addSlice(s *[]int, t int) {
	fmt.Printf("slice_addr: %p, cap: %d, value: %v\n", &s, cap(*s), *s)
	*s = append(*s, t)
	fmt.Printf("slice_addr: %p, cap: %d, value: %v\n", &s, cap(*s), *s)
}
