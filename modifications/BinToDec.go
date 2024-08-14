package modifications

import (
	"fmt"
	"strconv"
)

func BinToDec(arr *[]string, index int) bool{
	decimal, err := strconv.ParseInt((*arr)[index], 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	(*arr)[index] = strconv.Itoa(int(decimal))
	index--
	return true
}