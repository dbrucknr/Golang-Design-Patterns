package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	Data map[string]any
}

func (app *Application) RenderTemplate(w http.ResponseWriter, t string, data *TemplateData) {
	var tmpl *template.Template

	// If we are using the template cache, try and extract the template from the map, stored in the receiver
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}
	// If the template is not found in the cache, build it from disk
	if tmpl == nil {
		newTemplate, err := app.BuildTemplateFromDisk(t)
		if err != nil {
			log.Println("Error building template from disk:", err)
			return
		}
		log.Print("Building template from disk")
		tmpl = newTemplate
	}
	if data == nil {
		data = &TemplateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, data); err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Application) BuildTemplateFromDisk(t string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.gohtml",
		"./templates/partials/header.partial.gohtml",
		"./templates/partials/footer.partial.gohtml",
		fmt.Sprintf("./templates/%s", t),
	}
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}
	app.templateMap[t] = tmpl
	return tmpl, nil
}
