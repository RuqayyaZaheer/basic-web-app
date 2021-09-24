package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, temp string) {

	tc, err := RenderTemplateTest()
	if err != nil {
		log.Fatal(err)
	}
	t, ok := tc[temp]
	if !ok {
		log.Fatal(err)

	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err = buf.WriteTo(w)

	if err != nil {
		fmt.Println("Error writeing template to browser", err)
	}

}

func RenderTemplateTest() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*page.tmpl")

	if err != nil {
		return myCache, err

	}
	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matchs, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err

		}
		if len(matchs) > 0 {
			ts, err := ts.ParseGlob("./templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}

			myCache[name] = ts
		}

	}

	return myCache, nil

}
