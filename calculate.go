package main

import "math"

type PlayStyle struct {
	normal  int64 // 1 or 2
	special int64 // 1 or 2 or 4
}

func Point2Time(point int64, p PlayStyle) (x int64, y int64) {
	a := float64(p.special)
	b := float64(p.normal)
	x := b * float64(point) / float64(150.0*a+320.0*a*b)
	y := 150.0 * a * float64(point) / (53.0*150.0*a + 53.0*320.0*a*b)

	return searchSolution(x, y, p)
}

func requireItem(x, y int64, p PlayStyle) bool {
	if 53*p.normal*y-150*p.special*x > 0 {
		return true
	}
	return false
}

func searchSolution(x, y float64, p PlayStyle) (x, y int64) {
	x_up := math.Ceil(x)
	x_down := math.Trunc(x)

	y_up := math.Ceil(y)

	if requireItem(x_down, y_up, p) {
		return x_down, y_up
	}
	for inf_roop := 0; ; {
		y_up += 1
		if requireItem(x_down, y_up, p) {
			return x_down, y_up
		}
		if requireItem(x_up, y_up-1, p) {
			return x_up, y_up - 1
		}
	}
}
