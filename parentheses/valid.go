package parentheses

// IsValid returns true if parentheses match
func IsValid(s string) bool {
	var st []rune
	for _, e := range s {
		if e == '{' || e == '(' || e == '[' {
			st = append(st, e)
		} else if e == '}' || e == ')' || e == ']' {
			if len(st) == 0 {
				return false
			}
			top := st[len(st)-1]
			if !match(top, e) {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	return len(st) == 0
}

func match(r1, r2 rune) bool {
	return r1 == '{' && r2 == '}' || r1 == '(' && r2 == ')' || r1 == '[' && r2 == ']'
}
