package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/msdeboe/bookings/pkg/config"
	"github.com/msdeboe/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates gets template cache
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		//DEV MODE
		tc, _ = CreateTemplateCache()
	}
	//get the template cache from the app config

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}
	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := t.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}
	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmplCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl from templates folder
	pages, err := filepath.Glob("../../templates/*.page.tmpl")
	log.Println(pages)
	if err != nil {
		return tmplCache, err
	}

	//range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)                    //gets file name
		ts, err := template.New(name).ParseFiles(page) //ts - Template set
		if err != nil {
			return tmplCache, err
		}
		//get all layout files
		matches, err := filepath.Glob("../../templates/*.layout.tmpl")
		if err != nil {
			return tmplCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../../templates/*.layout.tmpl")
			if err != nil {
				return tmplCache, err
			}
		}

		tmplCache[name] = ts
	}

	return tmplCache, nil
}
