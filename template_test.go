package i18n

import (
	"bytes"
	"testing"
	"text/template"
)

func TestTemplate(t *testing.T) {
	parse, err := template.New("").Delims("{{", "}}").Funcs(template.FuncMap{}).Parse("hello {{._0}}")
	if err != nil {
		t.Error(err)
	}
	var buf bytes.Buffer
	param := make(map[string]interface{})
	param["_0"] = "world"
	err = parse.Execute(&buf, param)
	if err != nil {
		t.Error(err)
	}
	t.Log(buf.String())

}
