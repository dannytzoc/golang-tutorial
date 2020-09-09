package main

import "fmt"

func chaneValue(str *string) {
	*str = "changed"
}
func changeValue2(str string) {
	str = "changed"
}
func main() {
	x := 7
	y := &x

	fmt.Println(x, *y)
	*y = 8
	fmt.Println(x, *y)
	tochange := "hello"
	// fmt.Println(tochange)
	// chaneValue(&tochange)
	// fmt.Println(tochange)
	fmt.Println(tochange)
	changeValue2(tochange)
	fmt.Println(tochange)

}
