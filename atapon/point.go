package atapon

import "fmt"

type Point struct {
	point     int64
	item      int64
	ref_point int64
}

func NewPoint(point, item, ref_point int64) Point {
	return Point{point, item, ref_point}
}

func (p Point) Output() {
	fmt.Print("point:", p.point, " item:", p.item, " 目標point:", p.ref_point)
}

