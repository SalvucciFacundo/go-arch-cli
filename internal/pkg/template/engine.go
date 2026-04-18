package template

import (
	"embed"
	"io"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

// Templates es el FS embebido que contiene todas las plantillas.
//go:embed all:templates/*
var TemplatesFS embed.FS

type Engine struct {
	fs embed.FS
}

func NewEngine() *Engine {
	return &Engine{
		fs: TemplatesFS,
	}
}

func (e *Engine) Render(wr io.Writer, templatePath string, data interface{}) error {
	funcMap := template.FuncMap{
		"now": func() string {
			return time.Now().Format("2006-01-02 15:04:05")
		},
		"lower": strings.ToLower,
		"upper": strings.ToUpper,
	}

	// Cargamos la plantilla desde el FS embebido
	t, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).ParseFS(e.fs, filepath.Join("templates", templatePath))
	if err != nil {
		return err
	}

	return t.Execute(wr, data)
}
