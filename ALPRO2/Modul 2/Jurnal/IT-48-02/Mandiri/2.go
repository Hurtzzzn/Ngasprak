package main

import (
	"fmt"
	"math"
)

func konversiDerajatkeRadian(T float64) float64 {
	return T * math.Pi / 180
}

func waktu(V, T float64) float64 {
	var t float64
	const g float64 = 9.8
	t = V * math.Sin(konversiDerajatkeRadian(T)) / g
	return t
}

func jarak(V, T float64) float64 {
	var x float64
	const g float64 = 9.8
	x = V * V * math.Sin(2*konversiDerajatkeRadian(T)) / g
	return x
}

func ketinggian(V, T float64) float64 {
	var y float64
	const g float64 = 9.8
	y = V * V * math.Sin(konversiDerajatkeRadian(T)) * math.Sin(konversiDerajatkeRadian(T)) / (2.0 * g)
	return y
}

func main() {
	var V, T float64
	var waktuPosMax, jarak_hor, ketinggian_max float64
	fmt.Scan(&V, &T)
	waktuPosMax = waktu(V, T)
	jarak_hor = jarak(V, T)
	ketinggian_max = ketinggian(V, T)
	fmt.Printf("%.2f %.2f %.2f\n", waktuPosMax, jarak_hor, ketinggian_max)
}
