package main

import (
	"fmt"
	"os"
	"strings"
	"goreloaded/files"
	"goreloaded/modifications"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return
	}
	source, err1 := os.Open(args[1])
	destination, err2 := os.Open(args[2])
	defer source.Close()
	defer destination.Close()
	defer fmt.Println("both files " + source.Name() + " and " + destination.Name() + " are closed.")

	if files.GetFileSize(source.Name()) == 0 {
		fmt.Println("the source file is empty")
		return
	}

	if err1 != nil || err2 != nil {
		fmt.Println("Error: Unable to open the file.")
		return
	} else {
		fmt.Println("both files " + source.Name() + " and " + destination.Name() + " are opened.")
	}
	sourceData, err := os.ReadFile(source.Name())
	if err != nil {
		fmt.Println("Error: Unable to read the file.")
		return
	} else {
		fmt.Println("text successfully read from \"" + source.Name() + "\".")
	}
	sourceDataSlice := strings.Fields(string(sourceData))
	for i := 0; i < len(sourceDataSlice); i++ {
		if i>=1 && strings.Contains(sourceDataSlice[i], "(") && modifications.PrefixNumber(sourceDataSlice[i+1]) <=i && !(strings.Contains(sourceDataSlice[i+1], "0")&&modifications.PrefixNumber(sourceDataSlice[i+1]) ==1){
		if strings.Contains(sourceDataSlice[i], "(hex") {
			if !modifications.HexToDec(&sourceDataSlice, i-1){
				return
			}
			if  modifications.PrefixNumber(sourceDataSlice[i+1])!=1 && strings.Contains(sourceDataSlice[i+1], ")"){
				fmt.Println(source.Name() + " contain a wrongly formatted method.")
				return
			}
		} else if strings.Contains(sourceDataSlice[i], "(bin") {
			if !modifications.BinToDec(&sourceDataSlice, i-1){
				return
			}
			if  modifications.PrefixNumber(sourceDataSlice[i+1])!=1 && strings.Contains(sourceDataSlice[i+1], ")"){
				fmt.Println(source.Name() + " contain a wrongly formatted method.")
				return
			}
		} else if strings.Contains(sourceDataSlice[i], "(up") {
			modifications.ToUpperCase(&sourceDataSlice, sourceDataSlice[i+1], i-1)
		} else if strings.Contains(sourceDataSlice[i], "(low") {
			modifications.ToLowerCase(&sourceDataSlice, sourceDataSlice[i+1], i-1)
		} else if strings.Contains(sourceDataSlice[i], "(cap") {
			modifications.Capitalize(&sourceDataSlice, sourceDataSlice[i+1], i-1)
		}
	}else if strings.Contains(sourceDataSlice[i], "("){
		fmt.Println(source.Name() + " contains unachievable method.")
		return
	}
}
	str := modifications.FixPunctuation(sourceDataSlice)
	content := []byte(str)
	err3 := os.WriteFile(destination.Name(), []byte(content), 0644)

	if err3 != nil {
		fmt.Println("Error writing file:", err)
		return
	} else {
		fmt.Println("text successfully written in \"" + destination.Name() + "\".")
	}

}
