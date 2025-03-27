package locale

import (
	"embed"

	"github.com/goccy/go-yaml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed active.*.yaml
var localeFS embed.FS

var bundle *i18n.Bundle

func init() {
	bundle = i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)
	_, _ = bundle.LoadMessageFileFS(localeFS, "active.ja.yaml")
}

func NewLocalizer(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, lang)
}
