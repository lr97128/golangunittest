package Mysplit

import "strings"

func Mysplit(s, sep string) (result []string) {
	Index := strings.Index(s, sep)
	for Index >= 0 {
		str := s[:Index]
		if str != "" {
			result = append(result, str)
		}
		s = s[(Index + len(sep)):]
		Index = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
