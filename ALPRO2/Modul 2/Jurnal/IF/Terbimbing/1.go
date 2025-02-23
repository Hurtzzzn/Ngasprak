package main

import "fmt"

func main() {

	type rectangle struct {
		length    float64
		width     float64
		perimeter float64
		area      float64
	}

	var r rectangle
	var n, i int

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&r.length)
		fmt.Scan(&r.width)
		r.perimeter = (2 * r.length) + (2 * r.width)
		r.area = r.length * r.width
		fmt.Println(r.perimeter, r.area)
	}
}
