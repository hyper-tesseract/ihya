package templating

import (
	"bytes"
	"fmt"
	"html/template"
	"ihya-web/app/entity"
	hyper "ihya-web/app/hyper"
	"log"
	"net/http"
)

const (
	TemplatePath = "../app/adapter/templating/templates"
	ThemesPath   = "../app/adapter/templating/templates/themes"
)

var IhyaWrapperData = struct {
	FooterText   string
	PortfolioUrl string
	Author       string
	PageData     entity.PageData
}{
	FooterText:   "Theme by: boxmodel.dev | Ihya by",
	PortfolioUrl: "http://muhammadazeem.com",
	Author:       "Muhammad Azeem",
}

type TemplateRenderer struct {
	templates template.Template
}

func (t *TemplateRenderer) RenderPage(
	w http.ResponseWriter,
	pageData entity.PageData) error {

	IhyaWrapperData.PageData = pageData

	err := t.templates.ExecuteTemplate(w, "ihya-web", IhyaWrapperData)

	if err != nil {
		msg := fmt.Errorf("error executing template: %s", err)
		log.Print(msg)
		return msg
	}

	return nil
}

func renderTemplateByName(t *template.Template) interface{} {
	return func(name string, data interface{}) (template.HTML, error) {
		buf := bytes.NewBuffer([]byte{})
		err := t.ExecuteTemplate(buf, name, data)
		ret := template.HTML(buf.String())
		return ret, err
	}
}

func NewTemplateRenderer() *TemplateRenderer {
	t := template.New("")

	t.Funcs(map[string]interface{}{
		"IncludeTemplate": renderTemplateByName(t),
	})

	t, err := t.ParseGlob(hyper.FormatString("%s/*.html", TemplatePath))
	fmt.Printf("%s", err)
	t, err = t.ParseGlob(hyper.FormatString("%s/*/*.html", ThemesPath))
	fmt.Printf("%s", err)
	return &TemplateRenderer{templates: *t}
}
