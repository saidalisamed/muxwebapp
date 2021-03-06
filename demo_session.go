package main

import (
	"net/http"

	"github.com/saidalisamed/muxwebapp/utils"

	"github.com/gorilla/context"
)

func (a *App) sessionSet(w http.ResponseWriter, r *http.Request) {
	// We need this to avoid memory leak
	defer context.Clear(r)

	session, _ := a.Store.Get(r, "login")
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	session.Save(r, w)

	utils.JSONResp(w, http.StatusOK, map[string]string{"result": "Session has been set"})
}

func (a *App) sessionGet(w http.ResponseWriter, r *http.Request) {
	// We need this to avoid memory leak
	defer context.Clear(r)

	session, _ := a.Store.Get(r, "login")
	foo := session.Values["foo"]

	utils.JSONResp(w, http.StatusOK, map[string]string{"foo": foo.(string)})
}
