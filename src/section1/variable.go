package main

import "fmt"

func main() {

	var total int = 100
	total = total * 2
	fmt.Println(total)

	var utotal uint = 100
	utotal = utotal * 3
	fmt.Println(utotal)

	const str string = "aaa"
	fmt.Println(str)

	var b bool = true
	fmt.Println(b)

	one, _ := hoge()
	fmt.Println(one)

	items := [...]string{"hoge", "huga"}
	for _, item := range items {
		fmt.Println(item)
	}
}

func hoge() (string, string) {
	return "1st", "2nd"
}
