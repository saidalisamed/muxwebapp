package main

import (
	"log"
	"net/http"
)

func (a *App) templateDemo(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"title": a.Cfg.App.Title,
	}

	if err := a.Templ.ExecuteTemplate(w, "template_demo", m); err != nil {
		log.Println(err)
	}
}
