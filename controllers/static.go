package controllers

import (
	"go-webdevelopment/views"
	"html/template"
	"net/http"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHttp(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler (tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct { //this is an inline struct: https://www.digitalocean.com/community/tutorials/defining-structs-in-go#inline-structs
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Yes, there is a free version of the software. ",
		},
		{
			Question: "How do I get the free version?",
			Answer:   "You can download the free version from our website.",
		},
		{
			Question: "How do can I get in touch with you?",
			Answer:   "Contact me on <a href='mailto:ghostbusters@gassma.nn'>ghostbusters@gassma.nn</a>",
		},
}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}