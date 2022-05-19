package controllers

import (
	"net/http"

	"github.com/fgabrielle1445/web_app/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHTTP(w, http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}