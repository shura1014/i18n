package i18n

import (
	"fmt"
	"text/template"
)

type Bundle struct {
	// 名称
	Name string
	// 语言
	Lang string
	// 翻译文件
	Messages map[string]TemplateExecute
	// 文件路径
	Path string

	pre     string
	suf     string
	funcMap template.FuncMap
}

func NewBundle(name, lang, path, pre, suf string, funcMap template.FuncMap) *Bundle {
	bundle := Bundle{
		Name:    name,
		Lang:    lang,
		Path:    path,
		pre:     pre,
		suf:     suf,
		funcMap: funcMap,
	}
	bundle.Messages = make(map[string]TemplateExecute)
	messages := bundle.LoadMessages(path)

	if messages != nil {
		for key, msg := range messages {
			_template := CreateTemplate(key, msg, bundle.pre, bundle.suf, bundle.funcMap)
			bundle.Messages[key] = _template
		}
	} else {
		fmt.Println("Load messages failed", path)
	}
	return &bundle
}

func (b *Bundle) LoadMessages(path string) map[string]string {
	return b.LoadMessagesFromFile(path)
}

// LoadMessagesFromFile 加载解析文件
func (b *Bundle) LoadMessagesFromFile(path string) map[string]string {
	for _, parse := range parses {
		if parse.Support(path) {
			return parse.Parse(path)
		}
	}
	return nil
}

// GetMessage 获取信息
func (b *Bundle) GetMessage(key string, args any) string {
	if _template, ok := b.Messages[key]; ok {
		msg, err := _template.Execute(args)
		if err != nil {
			fmt.Println(err.Error())
			return ""
		}
		return msg
	}
	return ""
}
