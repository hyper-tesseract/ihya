package template

import (
	"html/template"
)

type TemplateLoader struct {
	T *template.Template
}

func NewTemplateLoader() *TemplateLoader {
	t := template.New("theme")
	t = template.Must(t.ParseGlob("templates/*.html"))
	return &TemplateLoader{T: t}
}
