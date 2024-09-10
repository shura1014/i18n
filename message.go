package i18n

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

// 定义i18n接口

type I18n interface {
	init()
	Translate(key string, args ...interface{}) string
	CustomTranslate(key string, params map[string]any) string
}

// DefaultI18n 默认实现
type DefaultI18n struct {
	bundle      map[string][]*Bundle
	baseDir     string
	paramPrefix string
}

func NewDefaultI18n(baseDir string) *DefaultI18n {
	d := &DefaultI18n{
		bundle:      make(map[string][]*Bundle),
		baseDir:     baseDir,
		paramPrefix: "p",
	}
	d.init()
	return d
}

// SetBaseDir 设置baseDir
func (d *DefaultI18n) SetBaseDir(baseDir string) {
	d.baseDir = baseDir
}

func (d *DefaultI18n) SetParamPrefix(prefix string) {
	d.paramPrefix = prefix
}

// 初始化
func (d *DefaultI18n) init() {
	d.LoadMessages(d.baseDir)
}

// LoadMessages 加载文件
func (d *DefaultI18n) LoadMessages(baseDir string) {
	// 判断是否是文件夹，读取文件夹下面的所有文件
	stat, err := os.Stat(baseDir)
	if err != nil {
		fmt.Println("Failed to read the file directory. dir is " + baseDir)
		return
	}
	if stat.IsDir() {
		files, err := os.ReadDir(baseDir)
		if err != nil {
			fmt.Println("Failed to read the file directory.")
			return
		}
		for _, file := range files {
			index := strings.LastIndex(file.Name(), ".")
			if index > 0 {
				// 约定文件名格式为：语言_国家.后缀
				lang := file.Name()[:index]
				bundle := NewBundle(file.Name(), lang, filepath.Join(baseDir, file.Name()), "{{", "}}", template.FuncMap{})
				if d.bundle[lang] == nil {
					bundles := make([]*Bundle, 0)
					d.bundle[lang] = bundles
				}
				d.bundle[lang] = append(d.bundle[lang], bundle)
			}
		}
	}
}

func (d *DefaultI18n) Translate(lang, key string, args ...interface{}) string {
	param := make(map[string]interface{})
	if args != nil {
		for index, arg := range args {
			param[d.paramPrefix+strconv.Itoa(index)] = arg
		}
	}
	return d.CustomTranslate(lang, key, param)
}

func (d *DefaultI18n) CustomTranslate(lang, key string, params map[string]any) string {
	// 翻译
	bundles := d.bundle[lang]
	if bundles != nil {
		for _, bundle := range bundles {
			s := bundle.GetMessage(key, params)
			if s != "" {
				return s
			}
		}
		fmt.Println("No international content for " + key + " was found ")
	}

	fmt.Println("bundle was not found according to the language " + lang)
	return key
}
