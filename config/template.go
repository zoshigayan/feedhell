package config

import (
	"github.com/labstack/echo/v4"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Renderer() *Template {
	tmpl := template.New("")
	err := filepath.Walk("views", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			if _, err = tmpl.ParseFiles(path); err != nil {
				log.Println(err)
			}
		}
		return err
	})
	if err != nil {
		panic(err)
	}

	return &Template{
		templates: tmpl,
	}
}
