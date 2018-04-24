package main

func Reverse(input string) string {
	s := strings.Fields(input)
	return strings.Join(reverse(s), " ")
}

func reverse(ss []string) []string {
	l := len(ss)
	if l <= 1 {
		return ss
	}
	for i := 0; i < l/2; i++ {
		ss[i], ss[l-1-i] = ss[l-1-i], ss[i]
	}
	return ss
}
