package routes_test

import (
	"html/template"
	"log"
	"path/filepath"
)

func getTemplate(name string) *template.Template {
	templateDir := "../templates/"

	templateMap := make(map[string]*template.Template)

	layouts, err := filepath.Glob(templateDir + "shared/*.html")
	if err != nil {
		log.Fatalf("test_helper: could not get layouts: %v\n", err)
	}

	templates, err := filepath.Glob(templateDir + "*.html")
	if err != nil {
		log.Fatalf("test_helper: could not get templates: %v\n", err)
	}

	for _, tpl := range templates {
		files := append(layouts, tpl)
		templateMap[filepath.Base(tpl)] = template.Must(template.ParseFiles(files...))
	}

	return templateMap[name]
}
