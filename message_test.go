package i18n

import "testing"

func TestMessages(t *testing.T) {
	i18n := NewDefaultI18n("message")
	langZh := "zh_cn"
	langEn := "en_us"
	en100001 := i18n.Translate(langEn, "100001", "shura")
	en100002 := i18n.Translate(langEn, "100002")
	zh100001 := i18n.Translate(langZh, "100001", "shura")
	zh100002 := i18n.Translate(langZh, "100002")
	param := make(map[string]interface{})
	param["p0"] = "word!!"
	customTranslate := i18n.CustomTranslate(langEn, "100001", param)
	t.Log(en100001)
	t.Log(en100002)
	t.Log(zh100001)
	t.Log(zh100002)
	t.Log(customTranslate)
}
