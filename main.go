package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)


func executeTemplate(w http.ResponseWriter, filepath string){
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)	
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

//Section 3 - Exercise 1 - Use URL Parameters
func galleryHandler(w http.ResponseWriter, r *http.Request) {
	imageID := chi.URLParam(r, "imageID")

	//fmt.Fprint(w, "<h1>Gallery Page </h1>")	
	w.Write([]byte(fmt.Sprintf("Image ID: %v", imageID)))
}


/* func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch(r.URL.Path) {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
} */


func main() {
	r := chi.NewRouter()
	//r.Use(middleware.Logger) //Section 3 - Exercise 2 Use Middleware Logger GLOBAL
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Route("/admin", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Get("/gallery/{imageID}", galleryHandler) //Section 3 - Exercise 1 - Use URL Parameters only for one route
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe("127.0.0.1:3000", r) //For windows - it needs the full path to avoid the Firewall question
}
