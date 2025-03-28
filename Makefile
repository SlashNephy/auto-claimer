build:
	go build ./cmd/auto-claimer

i18n-extract:
	go tool goi18n extract -sourceLanguage en -outdir ./locale -format yaml
	go tool goi18n merge -outdir ./locale -format yaml ./locale/active.en.yaml ./locale/active.ja.yaml

i18n-merge:
	go tool goi18n merge -outdir ./locale -format yaml ./locale/active.en.yaml ./locale/active.ja.yaml ./locale/translate.ja.yaml
