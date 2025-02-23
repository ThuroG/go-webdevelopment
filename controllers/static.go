package controllers

import (
	"go-webdevelopment/views"
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