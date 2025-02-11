package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprintln(w, "<h1>Hello, World!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page </h1> <p> To get in touch, please send an email to <a href=\"mailto:arthur.gassmann@outlook.com\">Arthur Gassmann</a></p>")	
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page </h1> 
	
	<p> Q: Is there a free Version </p>
	<p> A: Yes, there is a free version of the software. </p>
	
	<p> Q: How do I get the free version? </p>
	<p> A: You can download the free version from our website. </p>
	
	`)	
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
	http.ListenAndServe(":3000", r)
}
