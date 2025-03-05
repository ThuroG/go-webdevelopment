package main

import (
	"fmt"
	"go-webdevelopment/controllers"
	"go-webdevelopment/templates"
	"go-webdevelopment/views"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

// Section 3 - Exercise 1 - Use URL Parameters
func galleryHandler(w http.ResponseWriter, r *http.Request) {
	imageID := chi.URLParam(r, "imageID")
	views.Must(views.ParseFS(templates.FS, "gallery.gohtml", "tailwind.gohtml")).Execute(w, imageID) // Section 6 - Exercise: Define another static handler
	//fmt.Fprint(w, "<h1>Gallery Page </h1>")
	//w.Write([]byte(fmt.Sprintf("Image ID: %v", imageID)))
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

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))


	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.Route("/gallery", func(r chi.Router) {
		r.Use(middleware.Logger)
	    r.Get("/{imageID}", galleryHandler) //Section 3 - Exercise 1 - Use URL Parameters only for one route
	})

	//Use a controller for creating a new user
	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		 "signup.gohtml", "tailwind.gohtml",
		 ))
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Server is running on http://localhost:3000")
	http.ListenAndServe("127.0.0.1:3000", r) //For windows - it needs the full path to avoid the Firewall question
}
