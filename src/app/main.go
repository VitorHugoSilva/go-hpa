package main

import "fmt"

func main() {
	x := 0.0001
	for i := 0; i <= 1; i++ {
		x = Calc(x)
	}
	fmt.Println("Code.education Rocks!")
}