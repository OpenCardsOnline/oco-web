package handlers

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

var MainHeaderTitle = "OpenCardsOnline"

type PageData struct {
	Header  map[string]interface{}
	Content string
}

func ParseTemplates() {
	tmpl = template.Must(template.ParseGlob("./views/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./views/components/*.html"))
	tmpl = template.Must(tmpl.ParseGlob("./views/pages/*.html"))
}

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Header: map[string]interface{}{
			"Title": MainHeaderTitle,
			"Extra": "Not Found",
		},
		Content: "Page Not Found!",
	}

	err := tmpl.ExecuteTemplate(w, "pages/home", data)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func ComingSoonPageHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Header: map[string]interface{}{
			"Title": MainHeaderTitle,
			"Extra": "Coming Soon",
		},
		Content: "",
	}

	err := tmpl.ExecuteTemplate(w, "pages/coming-soon", data)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Header: map[string]interface{}{
			"Title": MainHeaderTitle,
			"Extra": "Welcome",
		},
		Content: "This is the home page",
	}

	err := tmpl.ExecuteTemplate(w, "pages/home", data)
	if err != nil {
		http.Error(w, "Unable to execute template", http.StatusInternalServerError)
	}
}
