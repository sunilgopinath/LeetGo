package lower

// ToLowerCase returns a lower case version of the string
func ToLowerCase(str string) string {
	ss := []rune(str)
	var rslt []rune

	for _, r := range ss {
		if r >= 65 && r <= 90 {
			r += (32)
		}
		rslt = append(rslt, r)
	}

	return string(rslt)
}
