package myfunctions

// Apostrophe function takes a slice of strings as input and handles cases where an apostrophe is present in the string.
func Apostrophe(s []string) []string {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == "'" && count == 0 {
			count++
			s[i+1] = "'" + s[i+1]
			s = append(s[:i], s[i+1:]...)
		} else if s[i] == "'" && count%2 == 1 {
			count++
			s[i-1] = s[i-1] + "'"
			s = append(s[:i], s[i+1:]...)
		} else if s[i] == "'" && count%2 == 0 {
			count++
			s[i+1] = "'" + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
