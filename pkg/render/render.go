package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(w http.ResponseWriter, temp string) {

	parsedTemplate, _ := template.ParseFiles("./templates/" + temp)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error to passing template", err)
		return
	}

}
