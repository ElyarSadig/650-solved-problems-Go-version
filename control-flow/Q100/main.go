package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print("$")
		}
		fmt.Println()
	}
}
