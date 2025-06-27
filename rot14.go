package piscine

func Rot14(input string) string {
	output := ""

	for _, x := range input {
		if 'A' <= x && x <= 'Z' {
			x = 'A' + (x-'A'+14)%26
		} else if 'a' <= x && x <= 'z' {
			x = 'a' + (x-'a'+14)%26
		}

		output += string(x)
	}

	return output
}