build:
	go build ./cmd/auto-claimer

translate-extract:
	go tool goi18n extract -outdir ./locale -format yaml

translate-merge:
	go tool goi18n merge -outdir ./locale -format yaml ./locale/active.en.yaml ./locale/translate.ja.yaml
