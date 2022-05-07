package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/volatiletech/authboss/v3"
)

func renderIndex(w http.ResponseWriter, r *http.Request) {
	mustRender(w, r, "index", authboss.HTMLData{})
}

func renderAccount(w http.ResponseWriter, r *http.Request) {
	mustRender(w, r, "account", authboss.HTMLData{})
}

func renderPrivacy(w http.ResponseWriter, r *http.Request) {
	mustRender(w, r, "privacy", authboss.HTMLData{})
}

func renderTerms(w http.ResponseWriter, r *http.Request) {
	mustRender(w, r, "terms", authboss.HTMLData{})
}

func mustRender(w http.ResponseWriter, r *http.Request, name string, data authboss.HTMLData) {
	// We've sort of hijacked the authboss mechanism for providing layout data
	// for our own purposes. There's nothing really wrong with this but it looks magical
	// so here's a comment.
	var current authboss.HTMLData
	dataIntf := r.Context().Value(authboss.CTXKeyData)
	if dataIntf == nil {
		current = authboss.HTMLData{}
	} else {
		current = dataIntf.(authboss.HTMLData)
	}

	current.MergeKV("csrf_token", nosurf.Token(r))
	current.MergeKV("user", r.Context().Value(authboss.CTXKeyUser))
	current.Merge(data)

	err := templates.Render(w, name, current)
	if err == nil {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = fmt.Fprintln(w, "Error occurred rendering template:", err)
}
