package main

import (
	"muxwebapp/utils"
	"net/http"
)

func (a *App) sessionSet(w http.ResponseWriter, r *http.Request) {
	session, _ := a.Store.Get(r, "login")
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	session.Save(r, w)

	utils.JSONResp(w, http.StatusOK, map[string]string{"result": "Session has been set"})
}

func (a *App) sessionGet(w http.ResponseWriter, r *http.Request) {
	session, _ := a.Store.Get(r, "login")
	foo := session.Values["foo"]

	utils.JSONResp(w, http.StatusOK, map[string]string{"foo": foo.(string)})
}
