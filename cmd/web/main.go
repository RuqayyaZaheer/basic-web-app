package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RuqayyaZaheer/basic-web-app/pkg/config"
	"github.com/RuqayyaZaheer/basic-web-app/pkg/handlers"
	"github.com/RuqayyaZaheer/basic-web-app/pkg/render"
)

func main() {

	var app config.AppConfig
	tc, err := render.RenderTemplateTest()
	if err != nil {
		log.Fatal("Can not create template cache")
	}
	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server is running in port:8000")

	_ = http.ListenAndServe(":8000", nil)
}
