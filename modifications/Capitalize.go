package modifications

func Capitalize(arr *[]string, method string, index int) {
	num := PrefixNumber(method)
	for i := 0; i < num; i++ {
		str := (*arr)[index]
		runes := []rune(str)
		for j := 0; j < len(runes); j++ {
			if j == 0 && runes[j] >= 'A' && runes[j] <= 'Z' {
				continue
			} else if j == 0 && runes[j] >= 'a' && runes[j] <= 'z' {
				runes[j] -= 32
			} else if j != 0 && runes[j] >= 'a' && runes[j] <= 'z' {
				continue
			} else if j != 0 && runes[j] >= 'A' && runes[j] <= 'Z' {
				runes[j] += 32
			}
		}
		(*arr)[index] = string(runes)
		index--
	}

}