package stringutil

// Overlapping String
func CommonStr(a, b string) string {
	if len(a) > len(b) {
		for i := 0; i < len(a)-len(b)+1; i++ {
			if a[i:i+len(b)] == b {
				return b

			}
		}
	}

	for i := 0; i < len(b)-len(a)+1; i++ {
		if b[i:i+len(b)] == a {
			return a
		}
	}

	return ""

}
