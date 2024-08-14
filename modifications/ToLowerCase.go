package modifications

func ToLowerCase(arr *[]string, method string, index int) {
	num := PrefixNumber(method)
	for i := 0; i < num; i++ {
		str := (*arr)[index]
		runes := []rune(str)
		for j := 0; j < len(runes); j++ {
			if runes[j] >= 'A' && runes[j] <= 'Z' {
				runes[j] += 32
			}
		}
		(*arr)[index] = string(runes)
		index--
	}
}