package main

import "fmt"

func main() {
	intslice := [10]int{0,1,2,3,4,5,6,7,8,9}

	for i := 0; i < 10; i++ {
		if intslice[i] % 2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i," is odd")
		}
	}
} 
