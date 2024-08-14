package modifications

import "strings"

func PrefixNumber(method string) int {
	num := 0
	if strings.Contains(method, ")") {
		for _, char := range method {
			if char > '0' && char <= '9' {
				num = num*10 + int(char-'0')
			}
		}
	}
	if num == 0 {
		return 1
	}
	return num
}