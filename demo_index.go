package main

import (
	"log"
	"net/http"
)

func (a *App) indexDemo(w http.ResponseWriter, r *http.Request) {
	m := map[string]interface{}{
		"title": a.Cfg.App.Title,
	}

	if err := a.Template.ExecuteTemplate(w, "demo_index", m); err != nil {
		log.Println(err)
	}
}
