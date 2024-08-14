package modifications

func ToUpperCase(arr *[]string, method string, index int) {
	num := PrefixNumber(method)
	for i := 0; i < num; i++ {
		str := (*arr)[index]
		runes := []rune(str)
		for j := 0; j < len(runes); j++ {
			if runes[j] >= 'a' && runes[j] <= 'z' {
				runes[j] -= 32
			}
		}
		(*arr)[index] = string(runes)
		index--
	}
}