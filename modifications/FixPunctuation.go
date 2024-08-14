package modifications

import (
	"strings"
)

func FixPunctuation(arr []string) string {
	str := ""
	for i := 0; i < len(arr); i++ {
		if strings.Contains(arr[i], "(") || strings.Contains(arr[i], ")") {
			continue
		} else if i != len(arr)-1 && (arr[i] == "a" || arr[i] == "A") {
			if i != 0 {
				str += " "
			}
			nextWord := arr[i+1]
			if nextWord[0] == 'a' || nextWord[0] == 'e' || nextWord[0] == 'i' || nextWord[0] == 'o' || nextWord[0] == 'u' || nextWord[0] == 'h' {
				str += arr[i] + "n"
			} else {
				str += arr[i]
			}
		} else {
			for j, char := range arr[i] {
				if str == "" {
					str += string(char)
				} else {
					if char == ',' || char == '?' || char == ';' || char == ':' || char == '!' {
						str += string(char)
					} else if char == '.' {
						if (str[len(str)-1] >= '0' && str[len(str)-1] <= '9') || (str[len(str)-1] >= 'a' && str[len(str)-1] <= 'z') || (str[len(str)-1] >= 'A' && str[len(str)-1] <= 'Z') || str[len(str)-1] == '.' {
							str += string(char)
						} else {
							str += " "
							str += string(char)
						}
					} else if char == '?' {
						if (str[len(str)-1] >= 'a' && str[len(str)-1] <= 'z') || (str[len(str)-1] >= 'A' && str[len(str)-1] <= 'Z') || str[len(str)-1] == '!' {
							str += string(char)
						} else {
							str += " "
							str += string(char)
						}
					} else if str[len(str)-1] == ',' || str[len(str)-1] == '?' || str[len(str)-1] == ';' || str[len(str)-1] == ':' {
						str += " "
						str += string(char)
					} else {
						if j == 0 && str[len(str)-1] != '\'' && char != '\'' {
							str += " "
						}
						str += string(char)
					}

				}
			}
		}
	}
	return str
}
