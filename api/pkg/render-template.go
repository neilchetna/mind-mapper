package pkg

import (
	"bytes"
	"text/template"
)

func RenderTemplate(templateContent string, data any) (string, error) {
	tmpl, err := template.New("tmpl").Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
