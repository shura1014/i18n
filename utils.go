package i18n

import (
	"bufio"
	"io"
	"os"
)

// ReadLine 按行读取文件
func ReadLine(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	var result []string
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err != nil && err != io.EOF {
			panic(err)
		}
		if err == io.EOF {
			break
		}

		result = append(result, string(line))

	}
	return result
}
