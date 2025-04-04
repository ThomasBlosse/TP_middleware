package config

import (
	"bytes"
	"embed"
	"github.com/adrg/frontmatter"
	"strings"
	"text/template"
)

//go:embed templates/*
var embeddedTemplates embed.FS

func GetStringFromEmbeddedTemplate(templatePath string, body interface{}) (content string, subject string, err error) {
	temp, err := template.ParseFS(embeddedTemplates, templatePath)
	if err != nil {
		return "", "", err
	}
	var tpl bytes.Buffer
	if err := temp.Execute(&tpl, body); err != nil {
		return "", "", err
	}

	var matter struct {
		Subject string `yaml:"subject"`
	}
	mailContent, err := frontmatter.Parse(strings.NewReader(tpl.String()), &matter)
	if err != nil {
		return "", "", err
	}

	return string(mailContent), matter.Subject, nil
}
