package main

import "fmt"

func main() {
	var v1 [3]int
	v1[0] = 10
	v1[1] = 20
	v1[2] = 30
	fmt.Println(len(v1), v1[0], v1[1], v1[2])

	v2 := [4]string{"a", "b", "c", "d"}
	fmt.Println(len(v2), v2[0], v2[1], v2[2], v2[3])

	v3 := [...]string{"aaa", "bbb"}
	// ここで要素を足す v3[2] = "ccc" はできない。宣言時に要素数は確定するようだ
	fmt.Println(len(v3), v3[0], v3[1])
}
