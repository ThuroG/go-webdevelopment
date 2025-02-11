package main

import (
	"fmt"
	"net/http"
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

type Router struct {}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch(r.URL.Path) {
		case "/":
			homeHandler(w, r)
		case "/contact":
			contactHandler(w, r)
		case "/faq":
			faqHandler(w, r)
		default:
			http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	var router Router
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":3000", router)
}
