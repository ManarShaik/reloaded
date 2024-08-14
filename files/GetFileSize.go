package files

import (
	"fmt"
	"os"
)

func GetFileSize(fileName string) int64 {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist")
		} else {
			fmt.Println("Error:", err)
		}
	}
	return fileInfo.Size()
}
