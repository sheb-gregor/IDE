// File generated by Gopher Sauce
// DO NOT EDIT!!
package handlers

import (
	methods "github.com/thestrukture/IDE/api/methods"
	"net/http"

	"github.com/gorilla/sessions"
)

//
func POSTApiGit(w http.ResponseWriter, r *http.Request, session *sessions.Session) (response string, callmet bool) {

	pkgName := r.FormValue("pkg")
	cmd := r.FormValue("cmd")

	if cmd == "commit" {
		mess := r.FormValue("message")
		response = mResponse(methods.CommitGit(pkgName, mess))
	}

	if cmd == "push" {
		methods.PushGit(pkgName)
		response = mResponse(false)
	}

	callmet = true
	return
}