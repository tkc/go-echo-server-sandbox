package template

import (
	"bytes"
	"encoding/json"
	"io"
	"text/template"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func GetTemplate() Template {

	funcMap := template.FuncMap{
		"json": func(data interface{}) string {
			jsonBytes, _ := json.Marshal(data)
			out := new(bytes.Buffer)
			json.Indent(out, jsonBytes, "", "")
			return out.String()
		},
	}

	t := &Template{
		templates: template.Must(template.New("").Funcs(funcMap).ParseGlob("public/views/*.html")),
	}

	return *t
}
