package piscine

func Unmatch(a []int) int {
	var count int = 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[i] == a[j] {
				count = count + 1
			}
		}
		if count%2 == 1 {
			return a[i]
		}
		count = 0
	}

	return -1
}