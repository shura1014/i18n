package i18n

import (
	"strings"
)

type PropertiesParse struct {
}

func (p *PropertiesParse) Support(path string) bool {
	return len(path) > 11 && path[len(path)-11:] == ".properties"
}

func (p *PropertiesParse) Parse(file string) map[string]string {
	lines := ReadLine(file)
	result := make(map[string]string)
	for _, line := range lines {
		index := strings.Index(line, "=")

		if index >= 0 {

			key := line[:index]
			msg := line[index+1:]
			result[key] = msg
		}

	}
	return result
}
