package handlers

import (
	"net/http"

	"github.com/RuqayyaZaheer/basic-web-app/pkg/render"
)

// Home is a function for showing home page
func Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl")

}

// About is a function for showing about page

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}
