// SQUARE

package main

import "fmt"

func main() {

	type Square struct {
		side      float64
		perimeter float64
		area      float64
	}

	var s Square
	var n, i int

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&s.side)
		s.perimeter = 4 * s.side
		s.area = s.side * s.side
		fmt.Println(s.perimeter, s.area)
	}
}
