package i18n

import (
	"bytes"
	"strings"
	"text/template"
)

// TemplateExecute 模版解析
// 两种情况，一种是需要模版解析，另一种不需要返回原始数据即可
type TemplateExecute interface {
	Execute(data any) (string, error)
}

// RowMessage 原始数据输出
type RowMessage struct {
	Message string
}

func (r *RowMessage) Execute(data any) (string, error) {
	return r.Message, nil
}

type TemplateMessage struct {
	RowMessage string
	Template   *template.Template
}

// NewTemplateMessage 创建TemplateMessage
func NewTemplateMessage(key, msg, pre, suf string, funcMap template.FuncMap) TemplateExecute {
	_template, err := template.New(key).Delims(pre, suf).Funcs(funcMap).Parse(msg)
	if err != nil {
		return nil

	}
	return &TemplateMessage{RowMessage: msg, Template: _template}
}

// Execute TemplateMessage
func (t *TemplateMessage) Execute(data any) (string, error) {
	var buf bytes.Buffer
	if err := t.Template.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// CreateTemplate 创建模版
func CreateTemplate(key, msg, pre, suf string, funcMap template.FuncMap) TemplateExecute {

	if pre == "" || suf == "" {
		return &RowMessage{Message: msg}
	}

	// 如果包含前缀，那么需要解析
	if strings.Contains(msg, pre) {
		return NewTemplateMessage(key, msg, pre, suf, funcMap)
	} else {
		return &RowMessage{Message: msg}
	}
}
