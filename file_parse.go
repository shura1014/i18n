// Package i18n 文件解析

package i18n

// 所有的解析实现
var parses []IParse

func init() {
	parses = append(parses, &JSONParse{})
	parses = append(parses, &PropertiesParse{})
}

// IParse 定义解析接口
type IParse interface {
	// Support 支持的类型
	Support(path string) bool
	// Parse 解析文件
	Parse(string) map[string]string
}

// RegisterParse 注册 Parse
func RegisterParse(parse IParse) {
	parses = append(parses, parse)
}
