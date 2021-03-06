package em_library

import (
	"context"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"path/filepath"
)

var (
	Instance_GoI18n *i18n.Bundle
)

// Initialization
// 初始化
// https://github.com/nicksnyder/go-i18n/tree/master/v2/example
func Init_GoI18n() {
	Instance_GoI18n = i18n.NewBundle(language.English)
	Instance_GoI18n.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	list, err := filepath.Glob("./storage/language/*.toml")
	if err != nil || len(list) < 1 {
		initLog.Fatalln("[ERROR]", "Failed to load language pack!", " Error:", err)
		return
	}

	for _, v:=range list {
		Instance_GoI18n.MustLoadMessageFile(v)
	}

	initLog.Println("[INFO]", "Successfully loaded Init_GoI18n.")
	return
}


type go_i18n struct {}

func NewGoI18n() *go_i18n {
	return &go_i18n{}
}


// Translate
// 翻译
func (this *go_i18n) TranslateString (s string, lang string) string {
	localizer := i18n.NewLocalizer(Instance_GoI18n, lang)
	ctx := localizer.MustLocalize(&i18n.LocalizeConfig{
		MessageID:      s,
	})
	return ctx
}


// Translate from request
// 从请求中翻译
func (this *go_i18n) TranslateFromRequest(ctx context.Context, str string) string {
	var lang string
	if ctx == nil {
		lang = "en"
	} else {
		v := ctx.Value("language")
		if v != nil && v != "" {
			lang = v.(string)
		} else {
			lang = "en"
		}
	}
	return this.TranslateString(str, lang)
}

