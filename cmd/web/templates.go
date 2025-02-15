package main

import (
	"html/template"
	"path/filepath"
	"time"

	"snippet-box.mms.io/internal/models"
)

type templateData struct {
	CurrentYear     int
	IsAuthenticated bool
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	CSRFToken       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	partials, err := filepath.Glob("./ui/html/partials/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		files := append(partials, "./ui/html/base.tmpl.html", page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
