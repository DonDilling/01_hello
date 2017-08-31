package main

import "fmt"

func main() {
	for i := 10000000; i < 10000010; i++ {
		fmt.Printf("%d \t %b \t %#X \n", i, i, i)
	}
}
