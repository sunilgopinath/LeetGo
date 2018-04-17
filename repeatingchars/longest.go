package repeatingchars

// LengthOfLongestSubstring returns length of longest unique char substring
func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	// try sliding windows and keep tmp map

	maxLen := 1 // must be at least one at this point
	for i := range s {
		maxLen = max(maxLen, unique(s[i:]))
	}
	return maxLen
}

func unique(s string) int {
	m := make(map[rune]int)
	for _, e := range s {
		if _, ok := m[e]; !ok {
			m[e] = 1
		} else {
			return len(m)
		}
	}
	return len(m)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
