package piscine

func Compact(ptr *[]string) int {
	x := 0
	var y []string
	for _, v := range *ptr {
		if v != "" && v != "false" && v != "0" {
			y = append(y, v)
			x++
		}
	}

	*ptr = y

	return x
}