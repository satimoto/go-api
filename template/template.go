package template

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
)

//go:embed html/*.html txt/*.txt
var templates embed.FS

func DefaultTemplate(templateName, extension string) (*template.Template, error) {
	return DefaultTemplateWithLocale(templateName, extension, "en")
}

func DefaultTemplateWithLocale(templateName, extension, locale string) (*template.Template, error) {
	t, err := template.ParseFS(templates, fmt.Sprintf("%s/%s.%s.%s", extension, templateName, locale, extension))

	if err != nil {
		t, err = template.ParseFS(templates, fmt.Sprintf("%s/%s.en.%s", extension, templateName, extension))
	}

	return t, err
}

func ParseEmailTemplate(templateName string, data interface{}) (string, string, error) {
	return ParseEmailTemplateWithLocale(templateName, "en", data)
}

func ParseEmailTemplateWithLocale(templateName, locale string, data interface{}) (string, string, error) {
	html, err := ParseTemplateWithLocale(templateName, "html", locale, data)

	if err != nil {
		return html, "", err
	}

	subject, err := ParseTemplateWithLocale(fmt.Sprintf("%s.subject", templateName), "txt", locale, data)

	return html, subject, err
}

func ParseTemplate(templateName, extension string, data interface{}) (string, error) {
	return ParseTemplateWithLocale(templateName, extension, "en", data)
}

func ParseTemplateWithLocale(templateName, extension, locale string, data interface{}) (string, error) {
	t, err := DefaultTemplateWithLocale(templateName, extension, locale)

	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
