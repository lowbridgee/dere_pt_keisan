package calculate

import "math"

type PlayStyle struct {
	Normal  int64 // 1 or 2
	Special int64 // 1 or 2 or 4
}

func Point2Time(point int64, p PlayStyle) (int64, int64) {
	a := float64(p.Special)
	b := float64(p.Normal)
	x := b * float64(point) / float64(150.0*a+320.0*a*b)
	y := 150.0 * a * float64(point) / (53.0*150.0*a + 53.0*320.0*a*b)

	return searchSolution(x, y, point, p)
}

func requireItem(x, y int64, p PlayStyle) bool {
	if 53*p.Normal*y-150*p.Special*x > 0 {
		return true
	}
	return false
}

func requirePoint(x,y,point int64, p PlayStyle) bool {
	if 53*y + 320*p.Special*x > point {
		return true
	}
		return false
}

func searchSolution(x, y float64, point int64, p PlayStyle) (int64, int64) {
	x_up := int64(math.Ceil(x))
	x_down := int64(math.Trunc(x))

	y_up := int64(math.Ceil(y))

	if ( requireItem(x_down, y_up, p) && requirePoint(x_down,y_up, point, p) ) {
		return x_down, y_up
	}
	for ;; {
		y_up += 1
		if (requireItem(x_down, y_up, p) && requirePoint(x_down,y_up, point, p) ){
			return x_down, y_up
		}
		if (requireItem(x_up, y_up-1, p) && requirePoint(x_down,y_up, point, p) ){
			return x_up, y_up - 1
		}
	}
}
