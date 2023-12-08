package main

import "fmt"

func main() {
	s1 := make([]int, 10, 20)
	s2 := append(s1, 1, 2, 3)

	var vs1 *int = &s1[2]
	var vs2 *int = &s2[2]

	equal := vs1 == vs2
	fmt.Println(vs1, vs2, equal) // 0xc0000a6010 0xc0000a6010 true

}
