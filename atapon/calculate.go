package atapon

import "fmt"

type Result struct {
	playtimes int64
}

type Speed struct {
	normal int64 // 1 or 2
	special int64 // 1 or 2 or 4
}

func NewSpeed(a , b int64) Speed {
	return Speed{a,b}
}

func item2point(p Point) int64 {
	return p.item / 150 * 320
}

func diff(p Point) int64 {
	return p.point - p.ref_point
}

func Playtimes(p Point) int {
	point := p.point
	for i := 1; ; i++ {
		point += 53
		if point > p.ref_point {
			return i
		}
	}
}

func Time2Point(point Point, conf Speed, second int64) int64 {
	playtimes := second / 150
	x := int64(53 * float64(conf.normal) / (150 * float64(conf.special) + 53 * float64(conf.normal)) * float64(playtimes))
	y := playtimes - x
	fmt.Print("x=",x,"  y=",y)
	return 320 * conf.special * x + 53 * y
}
