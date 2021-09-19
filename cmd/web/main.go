package main

import (
	"fmt"
	"net/http"

	"github.com/RuqayyaZaheer/basic-web-app/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server is running in port:8000")

	_ = http.ListenAndServe(":8000", nil)
}
