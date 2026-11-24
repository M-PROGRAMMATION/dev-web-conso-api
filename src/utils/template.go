package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

type Templates struct {
	templates *template.Template
}

func LoadTemplates(files ...string) (*Templates, error) {
	templates, err := template.ParseFiles(files...)
	if err != nil {
		return nil, fmt.Errorf("charger les templates HTML: %w", err)
	}

	return &Templates{templates: templates}, nil
}

func (t *Templates) Render(w http.ResponseWriter, name string, data any) error {
	var buffer bytes.Buffer
	if err := t.templates.ExecuteTemplate(&buffer, name, data); err != nil {
		return fmt.Errorf("afficher le template %q: %w", name, err)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := buffer.WriteTo(w); err != nil {
		return fmt.Errorf("écrire la réponse HTML: %w", err)
	}

	return nil
}
