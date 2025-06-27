package piscine

func ActiveBits(n int) int {
	x := 0
	y := 1

	for y*2 <= n {
		y *= 2
	}

	for y >= 1 {
		if n/y == 1 {
			x++
		}
		n = n % y
		y = y / 2
	}
	return x
}