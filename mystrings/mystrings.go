package mystrings

// names starting with capital letter will be visible in another packages
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// names starting with small letter will not be visible in another packages
func fun() string {
	return "this is fun"
}
