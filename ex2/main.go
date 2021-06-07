package main

import "fmt"

func main() {

	tr  := triangle{
		base: 5.0,
		height: 4.0,
	};

	var sq square
	sq.sideLength = 6.6
	fmt.Println(tr)
	fmt.Println(sq)
	fmt.Println(fmt.Sprintf("%.2f",sq.getArea()))
	fmt.Println(tr.getArea())
}
