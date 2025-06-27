package piscine

func Max(a []int) int {
	m := a[0]
	for _, v := range a {
		if v > m {
			m = v
		}
	}

	return m
}