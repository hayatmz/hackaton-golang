package piscine

func Join(strs []string, sep string) string {
	var x string

	for i := 0; i < len(strs); i++ {
		if i >= len(strs)-1 {
			x = x + strs[i]
		} else {
			x = x + strs[i] + sep
		}
	}
	return x
}