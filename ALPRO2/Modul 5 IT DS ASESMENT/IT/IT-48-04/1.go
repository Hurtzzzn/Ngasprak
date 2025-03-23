package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2, x3, y3 float64
	var g1, g2, g3, max_g float64

	fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)

	panjang_garis(x1, y1, x2, y2, x3, y3, &g1, &g2, &g3)

	max_garis(g1, g2, g3, &max_g)

	fmt.Println(max_g)
}

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func panjang_garis(a, b, c, d, e, f float64, g1, g2, g3 *float64) {
	*g1 = jarak(a, b, c, d)
	*g2 = jarak(c, d, e, f)
	*g3 = jarak(a, b, e, f)
}

func max_garis(a, b, c float64, max *float64) {
	*max = a
	if b > *max {
		*max = b
	}
	if c > *max {
		*max = c
	}
}
