package routes

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Cargar la plantilla HTML
	tmpl := template.Must(template.ParseFiles("template/index.html"))

	// Renderizar la plantilla con datos
	err := tmpl.Execute(w, nil)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}
