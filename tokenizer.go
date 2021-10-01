package go_tf_idf

import "strings"

func Tokenize(s string) []string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, ".", " ")
	s = strings.ReplaceAll(s, ",", " ")
	s = strings.ReplaceAll(s, ":", " ")
	s = strings.ReplaceAll(s, ";", " ")
	s = strings.ReplaceAll(s, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "/", " ")
	return strings.Fields(s)
}
