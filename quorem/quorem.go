package quorem

// patient 1
func QuoRem1(x, y int) (quo, rem int) {
	quo = x / y
	rem = x % y
	return
}

// patient 2
func QuoRem2(x, y int) (quo, rem int) {
	quo = x / y
	rem = x - quo*y
	return
}
