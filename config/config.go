package config

import (
	"bytes"
	"sync"
	"text/template"

	"github.com/starudream/go-lib/log"
)

type Config struct {
	Addr    string
	Storage string
	Owner   string
	Repo    string
	Name    string
	Proxy   string
}

var (
	C = &Config{}

	tplURL     *template.Template
	tplURLOnce sync.Once
)

func Default() {
	C = &Config{
		Addr:    "0.0.0.0:80",
		Storage: "/tmp",
		Owner:   "starudream",
		Repo:    "secret-tunnel",
		Name:    "{{.repo}}-{{.os}}-{{.arch}}-{{.platform}}-{{.ver}}.zip",
	}
}

func FormatURL(data map[string]any) (string, error) {
	tplURLOnce.Do(func() {
		text := "https://github.com/{{.owner}}/{{.repo}}/releases/download/{{.ver}}/" + C.Name

		tpl, err := template.New("tpl").Option("missingkey=error").Parse(text)
		if err != nil {
			log.Fatal().Msgf("parse template fail: %v", err)
		}

		tplURL = tpl
	})

	m := map[string]any{
		"owner": C.Owner,
		"repo":  C.Repo,
	}

	for k, v := range data {
		m[k] = v
	}

	bb := &bytes.Buffer{}
	err := tplURL.Execute(bb, m)
	return bb.String(), err
}
