package i18n

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON解析实现

type JSONParse struct {
}

// Support Json实现
func (j *JSONParse) Support(path string) bool {
	return len(path) > 5 && path[len(path)-5:] == ".json"
}

func (j *JSONParse) Parse(file string) map[string]string {
	return ParseJSON(file)
}

func ParseJSON(file string) map[string]string {
	// 解析json文件
	readFile, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		panic(err)
		return nil
	}

	// 实例化一个map
	bundle := make(map[string]string)
	err = json.Unmarshal(readFile, &bundle)
	if err != nil {
		panic(err)
		return nil
	}
	return bundle
}
